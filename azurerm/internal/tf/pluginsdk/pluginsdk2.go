package pluginsdk

/*
This is tracking the changes necessary for Plugin SDKv2

1. Base Layer
2. Package validation
	All validation functions need to be moved into `validation` - since otherwise the test namespace needs fixing
    Should we update resources to be Typed?
2. Updating the Customize Diff functions to use the new wrapper
3. Updating the Import functions to use the new wrappers
4. Gradually updating each service package to use the new wrappers
	NOTE: validation functions may need converting over to return the function rather than the type
5. Upgrade to Plugin SDKv2
*/