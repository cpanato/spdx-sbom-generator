// SPDX-License-Identifier: Apache-2.0

package models

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

// IPlugin ...
type IPlugin interface {
	SetRootModule(path string) error
	GetVersion() (string, error)
	GetMetadata() PluginMetadata
	GetRootModule(path string) (*Module, error)
	ListUsedModules(path string) ([]Module, error)
	ListModulesWithDeps(path string) ([]Module, error)
	IsValid(path string) bool
	HasModulesInstalled(path string) error
}

// PluginMetadata ...
type PluginMetadata struct {
	Name       string
	Slug       string
	Manifest   []string
	ModulePath []string
}

// Module ... ...
type Module struct {
	Version          string `json:"Version,omitempty"`
	Name             string
	Path             string `json:"Path,omitempty"`
	LocalPath        string `json:"Dir,noempty"`
	Supplier         SupplierContact
	PackageURL       string
	CheckSum         *CheckSum
	PackageHomePage  string
	LicenseConcluded string
	LicenseDeclared  string
	CommentsLicense  string
	OtherLicense     []*License
	Copyright        string
	PackageComment   string
	Root             bool
	Modules          map[string]*Module
}

// SupplierContact ...
type SupplierContact struct {
	Type  TypeContact
	Name  string
	Email string
}

// TypeContact ...
type TypeContact string

const (
	Person       TypeContact = "Person"
	Organization TypeContact = "Organization"
)

type CheckSum struct {
	Algorithm HashAlgorithm
	Content   []byte
	Value     string
}

func (c *CheckSum) String() string {
	if c.Value == "" {
		c.Value = c.calculateCheckSum(c.Content)
	}

	return fmt.Sprintf("%v: %s", c.Algorithm, c.Value)
}

func (c *CheckSum) calculateCheckSum(content []byte) string {
	var h hash.Hash
	switch c.Algorithm {
	case HashAlgoSHA256:
		h = sha256.New()
	case HashAlgoSHA512:
		h = sha512.New()
	default:
		h = sha1.New()
	}
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}

// HashAlgorithm ...
type HashAlgorithm string

const (
	HashAlgoSHA1   HashAlgorithm = "SHA1"
	HashAlgoSHA224 HashAlgorithm = "SHA224"
	HashAlgoSHA256 HashAlgorithm = "SHA256"
	HashAlgoSHA384 HashAlgorithm = "SHA384"
	HashAlgoSHA512 HashAlgorithm = "SHA512"
	HashAlgoMD2    HashAlgorithm = "MD2"
	HashAlgoMD4    HashAlgorithm = "MD4"
	HashAlgoMD5    HashAlgorithm = "MD5"
	HashAlgoMD6    HashAlgorithm = "MD6"
)

// License ...
type License struct {
	ID            string
	Name          string
	ExtractedText string
	Comments      string
	File          string
}
