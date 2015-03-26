var React = require('react');

var IsolationActions  = require('../actions/IsolationActions')
var IsolationStore = require('../stores/IsolationStore')
var DepStore = require('../stores/DepStore')
var IsolationList = require("./IsolationList.jsx")
var IsolationForm = require('./IsolationForm.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: IsolationStore.state(),
      depData: DepStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	IsolationStore.on(IsolationStore.CHANGED, this.onStoreChange)
    DepStore.on(DepStore.CHANGED, this.onStoreChange)
    IsolationActions.refresh()
  },

  componentDidUnmount: function() {
  	IsolationStore.removeListener(IsolationStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({data: IsolationStore.state(), depData: DepStore.state()})
  },

  openIsolationForm: function() {
    this.setState({ showForm: true });
  },

  closeIsolationForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    return <div style={{position: 'relative'}}>
      <h2 className="ui top attached header">
        <div className="content">
          Isolations
          <div style={{paddingRight: "180px"}} className="sub header">Isolations make groups of dependencies in specific states accessible under a certain name</div>
        </div>
      </h2>

      <button style={{position: 'absolute', top: '18px', right: '10px'}} className="ui secondary labeled icon button" onClick={this.openIsolationForm}><i className="plus icon"></i>New Isolation</button>
      
      { this.state.showForm ? <IsolationForm closeFormFn={this.closeIsolationForm} isolation={this.state.data.get('selection')}/> : null }
   
      <div className="ui bottom attached segment"><IsolationList deps={this.state.depData.get('deps')} selection={this.state.data.get('selection')} isolations={this.state.data.get('isolations')}/></div>
    </div>;
  }
});