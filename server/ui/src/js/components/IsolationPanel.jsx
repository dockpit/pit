var React = require('react');

var IsolationActions  = require('../actions/IsolationActions')
var IsolationStore = require('../stores/IsolationStore')
var IsolationList = require("./IsolationList.jsx")
var IsolationForm = require('./IsolationForm.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: IsolationStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	IsolationStore.on(IsolationStore.CHANGED, this.onStoreChange)
    IsolationActions.refresh()
  },

  componentDidUnmount: function() {
  	IsolationStore.removeListener(IsolationStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({data: IsolationStore.state()})
  },

  openIsolationForm: function() {
    this.setState({ showForm: true });
  },

  closeIsolationForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    return <div>
      <h2>Isolations <button className="dp-add" onClick={this.openIsolationForm}>+ New Isolation</button></h2>
      { this.state.showForm ? <IsolationForm closeFormFn={this.closeIsolationForm} isolation={this.state.data.get('selection')}/> : null }
      
      <IsolationList selection={this.state.data.get('selection')} isolations={this.state.data.get('isolations')}/>      
    </div>;
  }
});