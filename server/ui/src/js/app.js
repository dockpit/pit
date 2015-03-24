var React = require('react');
var Dispatcher = require('flux').Dispatcher;

var IsolationPanel = require("./components/IsolationPanel.jsx");
var DepPanel = require("./components/DepPanel.jsx");

//render panels on pages that contain them
//replace this later for a true one-page-app
var is = document.getElementById("isolation-section")
if (is) {
  React.render(<IsolationPanel/>, is)
}

var ds = document.getElementById("dep-section")
if (ds) {
  React.render(<DepPanel/>, ds)
}