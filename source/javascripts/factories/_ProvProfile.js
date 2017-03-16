(function() {

"use strict";

angular.module("BitriseWorkflowEditor").factory("ProvProfile", function() {

	var ProvProfile = function() {
		this.name;
	};

	ProvProfile.prototype.download = function() {
		window.location = this.name;
	};

	return ProvProfile;

});

})();
