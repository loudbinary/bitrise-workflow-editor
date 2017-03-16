(function() {

"use strict";

angular.module("BitriseWorkflowEditor").factory("GenericFile", function() {

	var GenericFile = function() {
		this.name;
	};

	GenericFile.prototype.download = function() {
		window.location = this.name;
	};

	return GenericFile;

});

})();
