var React = require('react');

var IsolationActions  = require('../actions/IsolationActions')
var IsolationStore = require('../stores/IsolationStore')
var IsolationList = require("./IsolationList.jsx")

module.exports = React.createClass({

  getInitialState: function() {
    return IsolationStore.state()
  },

  componentDidMount: function() {
  	IsolationStore.on(IsolationStore.CHANGED, this.onChange)
    IsolationActions.refresh()
  },

  componentDidUnmount: function() {
  	IsolationStore.removeListener(IsolationStore.CHANGED, this.onChange)
  },

  onChange: function() {
  	this.setState(IsolationStore.state())
  },

  render: function() {
    return <div>
      <h2>Isolations</h2>
      <IsolationList/>
    </div>;
  }
});