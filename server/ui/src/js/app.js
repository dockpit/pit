var React = require('react');
var Dispatcher = require('flux').Dispatcher;

var IsolationPanel = require("./components/IsolationPanel.jsx");
var DepPanel = require("./components/DepPanel.jsx");
var EditorWorkspace = require("./components/EditorWorkspace.jsx");

//render panels on pages that contain them
//replace this later for a true one-page-app
var is = document.getElementById("dp-iso-panel")
if (is) {
  React.render(<IsolationPanel/>, is)
}

// path: /
var ds = document.getElementById("dp-dep-panel")
if (ds) {
  React.render(<DepPanel/>, ds)
}

// path: /deps/:dname/states/:sname
var ew = document.getElementById("dp-editor-workspace")
if (ew) {
  var parts = window.location.pathname.split('/')
  React.render(<EditorWorkspace depName={parts[2]} stateName={parts[4]}/>, ew)
}