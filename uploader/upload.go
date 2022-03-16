package uploader

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
	"strings"
)

func Init(githubToken, ownerName, repoName, sarifContent string, prNumber int) *Uploader {
	return &Uploader{
		GithubToken:  githubToken,
		OwnerName:    ownerName,
		RepoName:     repoName,
		SarifContent: sarifContent,
		PRNumber:     prNumber,
	}
}

func (u *Uploader) UploadSARIF() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: u.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pr, _, err := client.PullRequests.Get(ctx, u.OwnerName, u.RepoName, u.PRNumber)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(u.SarifContent)); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	sarifGzipped := base64.StdEncoding.EncodeToString(b.Bytes())
	ref := fmt.Sprintf("refs/pull/%v/head", u.PRNumber)
	_, _, err = client.CodeScanning.UploadSarif(ctx, u.OwnerName, u.RepoName, &github.SarifAnalysis{
		CommitSHA: pr.Head.SHA,
		Ref:       &ref,
		Sarif:     &sarifGzipped,
	})

	if strings.Contains(err.Error(), "job scheduled") {
		return nil
	} else {
		return err
	}
}
