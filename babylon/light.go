package babylon

// SetIntensity calls the JavaScript method of the same name.
func (l *Light) SetIntensity(intensity float64) {
	l.p.Set("intensity", intensity)
}

// ExcludedMeshes returns the ExcludedMeshes property of class Light.
//
// https://doc.babylonjs.com/api/classes/babylon.light#excludedmeshes
func (l *Light) ExcludedMeshes() *AbstractMeshArray {
	retVal := l.p.Get("excludedMeshes")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), l.ctx))
	}
	return &AbstractMeshArray{p: retVal, am: result}
}

// SetExcludedMeshes sets the ExcludedMeshes property of class Light.
//
// https://doc.babylonjs.com/api/classes/babylon.light#excludedmeshes
func (l *Light) SetExcludedMeshes(excludedMeshes []*AbstractMesh) *Light {
	l.p.Set("excludedMeshes", excludedMeshes) // TODO: fix this.
	return l
}

// IncludedOnlyMeshes returns the IncludedOnlyMeshes property of class Light.
//
// https://doc.babylonjs.com/api/classes/babylon.light#includedonlymeshes
func (l *Light) IncludedOnlyMeshes() *AbstractMeshArray {
	retVal := l.p.Get("includedOnlyMeshes")
	result := []*AbstractMesh{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AbstractMeshFromJSObject(retVal.Index(ri), l.ctx))
	}
	return &AbstractMeshArray{p: retVal, am: result}
}

// SetIncludedOnlyMeshes sets the IncludedOnlyMeshes property of class Light.
//
// https://doc.babylonjs.com/api/classes/babylon.light#includedonlymeshes
func (l *Light) SetIncludedOnlyMeshes(includedOnlyMeshes []*AbstractMesh) *Light {
	l.p.Set("includedOnlyMeshes", includedOnlyMeshes) // TODO: fix this.
	return l
}
