// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SkyMaterial represents a babylon.js SkyMaterial.
// This is the sky material which allows to create dynamic and texture free effects for skyboxes.
//
// See: https://doc.babylonjs.com/extensions/sky
type SkyMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SkyMaterial) JSObject() js.Value { return s.p }

// SkyMaterial returns a SkyMaterial JavaScript class.
func (ba *Babylon) SkyMaterial() *SkyMaterial {
	p := ba.ctx.Get("SkyMaterial")
	return SkyMaterialFromJSObject(p, ba.ctx)
}

// SkyMaterialFromJSObject returns a wrapped SkyMaterial JavaScript class.
func SkyMaterialFromJSObject(p js.Value, ctx js.Value) *SkyMaterial {
	return &SkyMaterial{p: p, ctx: ctx}
}

// SkyMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func SkyMaterialArrayToJSArray(array []*SkyMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSkyMaterial returns a new SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#constructor
func (ba *Babylon) NewSkyMaterial(name string, scene *Scene) *SkyMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("SkyMaterial").New(args...)
	return SkyMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#bindforsubmesh
func (s *SkyMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	s.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#clone
func (s *SkyMaterial) Clone(name string) *SkyMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("clone", args...)
	return SkyMaterialFromJSObject(retVal, s.ctx)
}

// SkyMaterialDisposeOpts contains optional parameters for SkyMaterial.Dispose.
type SkyMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#dispose
func (s *SkyMaterial) Dispose(opts *SkyMaterialDisposeOpts) {
	if opts == nil {
		opts = &SkyMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	s.p.Call("dispose", args...)
}

// GetAlphaTestTexture calls the GetAlphaTestTexture method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#getalphatesttexture
func (s *SkyMaterial) GetAlphaTestTexture() *BaseTexture {

	retVal := s.p.Call("getAlphaTestTexture")
	return BaseTextureFromJSObject(retVal, s.ctx)
}

// GetAnimatables calls the GetAnimatables method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#getanimatables
func (s *SkyMaterial) GetAnimatables() []*IAnimatable {

	retVal := s.p.Call("getAnimatables")
	result := []*IAnimatable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IAnimatableFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// GetClassName calls the GetClassName method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#getclassname
func (s *SkyMaterial) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// SkyMaterialIsReadyForSubMeshOpts contains optional parameters for SkyMaterial.IsReadyForSubMesh.
type SkyMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#isreadyforsubmesh
func (s *SkyMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *SkyMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &SkyMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := s.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#needalphablending
func (s *SkyMaterial) NeedAlphaBlending() bool {

	retVal := s.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#needalphatesting
func (s *SkyMaterial) NeedAlphaTesting() bool {

	retVal := s.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Parse calls the Parse method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#parse
func (s *SkyMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *SkyMaterial {

	args := make([]interface{}, 0, 3+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := s.p.Call("Parse", args...)
	return SkyMaterialFromJSObject(retVal, s.ctx)
}

// Serialize calls the Serialize method on the SkyMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#serialize
func (s *SkyMaterial) Serialize() js.Value {

	retVal := s.p.Call("serialize")
	return retVal
}

// Azimuth returns the Azimuth property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#azimuth
func (s *SkyMaterial) Azimuth() float64 {
	retVal := s.p.Get("azimuth")
	return retVal.Float()
}

// SetAzimuth sets the Azimuth property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#azimuth
func (s *SkyMaterial) SetAzimuth(azimuth float64) *SkyMaterial {
	s.p.Set("azimuth", azimuth)
	return s
}

// CameraOffset returns the CameraOffset property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#cameraoffset
func (s *SkyMaterial) CameraOffset() *Vector3 {
	retVal := s.p.Get("cameraOffset")
	return Vector3FromJSObject(retVal, s.ctx)
}

// SetCameraOffset sets the CameraOffset property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#cameraoffset
func (s *SkyMaterial) SetCameraOffset(cameraOffset *Vector3) *SkyMaterial {
	s.p.Set("cameraOffset", cameraOffset.JSObject())
	return s
}

// Distance returns the Distance property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#distance
func (s *SkyMaterial) Distance() float64 {
	retVal := s.p.Get("distance")
	return retVal.Float()
}

// SetDistance sets the Distance property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#distance
func (s *SkyMaterial) SetDistance(distance float64) *SkyMaterial {
	s.p.Set("distance", distance)
	return s
}

// Inclination returns the Inclination property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#inclination
func (s *SkyMaterial) Inclination() float64 {
	retVal := s.p.Get("inclination")
	return retVal.Float()
}

// SetInclination sets the Inclination property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#inclination
func (s *SkyMaterial) SetInclination(inclination float64) *SkyMaterial {
	s.p.Set("inclination", inclination)
	return s
}

// Luminance returns the Luminance property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#luminance
func (s *SkyMaterial) Luminance() float64 {
	retVal := s.p.Get("luminance")
	return retVal.Float()
}

// SetLuminance sets the Luminance property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#luminance
func (s *SkyMaterial) SetLuminance(luminance float64) *SkyMaterial {
	s.p.Set("luminance", luminance)
	return s
}

// MieCoefficient returns the MieCoefficient property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#miecoefficient
func (s *SkyMaterial) MieCoefficient() float64 {
	retVal := s.p.Get("mieCoefficient")
	return retVal.Float()
}

// SetMieCoefficient sets the MieCoefficient property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#miecoefficient
func (s *SkyMaterial) SetMieCoefficient(mieCoefficient float64) *SkyMaterial {
	s.p.Set("mieCoefficient", mieCoefficient)
	return s
}

// MieDirectionalG returns the MieDirectionalG property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#miedirectionalg
func (s *SkyMaterial) MieDirectionalG() float64 {
	retVal := s.p.Get("mieDirectionalG")
	return retVal.Float()
}

// SetMieDirectionalG sets the MieDirectionalG property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#miedirectionalg
func (s *SkyMaterial) SetMieDirectionalG(mieDirectionalG float64) *SkyMaterial {
	s.p.Set("mieDirectionalG", mieDirectionalG)
	return s
}

// Rayleigh returns the Rayleigh property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#rayleigh
func (s *SkyMaterial) Rayleigh() float64 {
	retVal := s.p.Get("rayleigh")
	return retVal.Float()
}

// SetRayleigh sets the Rayleigh property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#rayleigh
func (s *SkyMaterial) SetRayleigh(rayleigh float64) *SkyMaterial {
	s.p.Set("rayleigh", rayleigh)
	return s
}

// SunPosition returns the SunPosition property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#sunposition
func (s *SkyMaterial) SunPosition() *Vector3 {
	retVal := s.p.Get("sunPosition")
	return Vector3FromJSObject(retVal, s.ctx)
}

// SetSunPosition sets the SunPosition property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#sunposition
func (s *SkyMaterial) SetSunPosition(sunPosition *Vector3) *SkyMaterial {
	s.p.Set("sunPosition", sunPosition.JSObject())
	return s
}

// Turbidity returns the Turbidity property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#turbidity
func (s *SkyMaterial) Turbidity() float64 {
	retVal := s.p.Get("turbidity")
	return retVal.Float()
}

// SetTurbidity sets the Turbidity property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#turbidity
func (s *SkyMaterial) SetTurbidity(turbidity float64) *SkyMaterial {
	s.p.Set("turbidity", turbidity)
	return s
}

// UseSunPosition returns the UseSunPosition property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#usesunposition
func (s *SkyMaterial) UseSunPosition() bool {
	retVal := s.p.Get("useSunPosition")
	return retVal.Bool()
}

// SetUseSunPosition sets the UseSunPosition property of class SkyMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.skymaterial#usesunposition
func (s *SkyMaterial) SetUseSunPosition(useSunPosition bool) *SkyMaterial {
	s.p.Set("useSunPosition", useSunPosition)
	return s
}
