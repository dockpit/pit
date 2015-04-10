var React = require('react');
var Dispatcher = require('./dispatcher/AppDispatcher')

var IsolationPanel = require("./components/IsolationPanel.jsx");
var DepPanel = require("./components/DepPanel.jsx");
var EditorWorkspace = require("./components/EditorWorkspace.jsx");

var DepActions = require('./actions/DepActions')
var IsolationActions = require('./actions/IsolationActions')
var EditorActions = require('./actions/EditorActions')

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
  React.render(<EditorWorkspace dockerHostIp={DOCKER_HOST_IP} depId={parts[2]} stateId={parts[4]}/>, ew)
}

//@todo find more elegant solution
window.depIconPath = function(dep, scale) {
  return tmplIconPath(dep.get('template'), scale)
}

window.tmplIconPath = function(tmpl, scale) {
 return tmpl.get('icons')+"/icon_t_"+tmpl.get('id') + (scale ? "@"+scale+"x" : "") + ".png" 
}

window.BRAND_BLUE = "#8CD2F4"

// track certain action
Dispatcher.register(function(a){
  switch (a.type) {
    case DepActions.CREATE:
      analytics.track('Created Dep');
      break

    case DepActions.ADD_DEP_STATE:
      analytics.track('Created Dep State');
      break

    case IsolationActions.CREATE:
      analytics.track('Created Isolation');
      break

    case IsolationActions.ADD_STATE:
      analytics.track('Added Dep State to Isolation');
      break

    case EditorActions.START_BUILD:
      analytics.track('Started Build');
      break
  }
})