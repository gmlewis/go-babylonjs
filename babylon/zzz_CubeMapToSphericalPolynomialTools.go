// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CubeMapToSphericalPolynomialTools represents a babylon.js CubeMapToSphericalPolynomialTools.
// Helper class dealing with the extraction of spherical polynomial dataArray
// from a cube map.
type CubeMapToSphericalPolynomialTools struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (c *CubeMapToSphericalPolynomialTools) JSObject() js.Value { return c.p }

// CubeMapToSphericalPolynomialTools returns a CubeMapToSphericalPolynomialTools JavaScript class.
func (b *Babylon) CubeMapToSphericalPolynomialTools() *CubeMapToSphericalPolynomialTools {
	p := b.ctx.Get("CubeMapToSphericalPolynomialTools")
	return CubeMapToSphericalPolynomialToolsFromJSObject(p)
}

// CubeMapToSphericalPolynomialToolsFromJSObject returns a wrapped CubeMapToSphericalPolynomialTools JavaScript class.
func CubeMapToSphericalPolynomialToolsFromJSObject(p js.Value) *CubeMapToSphericalPolynomialTools {
	return &CubeMapToSphericalPolynomialTools{p: p}
}

// TODO: methods
