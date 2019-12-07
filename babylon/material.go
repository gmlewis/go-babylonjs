package babylon

// DiffuseTexture returns the DiffuseTexture property of class Material.
//
// https://doc.babylonjs.com/api/classes/babylon.standardmaterial#diffusetexture
func (s *Material) DiffuseTexture() *BaseTexture {
	retVal := s.p.Get("diffuseTexture")
	return BaseTextureFromJSObject(retVal, s.ctx)
}

// SetDiffuseTexture sets the DiffuseTexture property of class Material.
//
// https://doc.babylonjs.com/api/classes/babylon.standardmaterial#diffusetexture
func (s *Material) SetDiffuseTexture(diffuseTexture *BaseTexture) *Material {
	s.p.Set("diffuseTexture", diffuseTexture.JSObject())
	return s
}

// SpecularColor returns the SpecularColor property of class Material.
//
// https://doc.babylonjs.com/api/classes/babylon.standardmaterial#specularcolor
func (s *Material) SpecularColor() *Color3 {
	retVal := s.p.Get("specularColor")
	return Color3FromJSObject(retVal, s.ctx)
}

// SetSpecularColor sets the SpecularColor property of class Material.
//
// https://doc.babylonjs.com/api/classes/babylon.standardmaterial#specularcolor
func (s *Material) SetSpecularColor(specularColor *Color3) *Material {
	s.p.Set("specularColor", specularColor.JSObject())
	return s
}
