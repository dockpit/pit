var React = require('react');

var DepActions  = require('../actions/DepActions')
var DepStore = require('../stores/DepStore')
var DepItem = require("./DepItem.jsx")
var DepForm = require('./DepForm.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: DepStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	DepStore.on(DepStore.CHANGED, this.onStoreChange)
    DepActions.refresh()
  },

  componentDidUnmount: function() {
  	DepStore.removeListener(DepStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({data: DepStore.state()})
  },

  openDepForm: function() {
    this.setState({ showForm: true });
  },

  closeDepForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    return <div style={{position: 'relative'}}>
      <h2 className="ui top attached header">
        <div className="content">
          Dependencies
          <div style={{paddingRight: '200px'}} className="sub header">{'Other services you project depends on: Databases, Queues, HTTP API\'s etc'}</div>
        </div>
      </h2>

      <button style={{position: 'absolute', top: '18px', right: '10px'}} className="ui labeled icon button primary" onClick={this.openDepForm}><i className="plus icon"></i>New Dependency</button>

      { this.state.showForm ? <DepForm closeFormFn={this.closeDepForm} dep={this.state.data.get('selection')}/> : null }

      {this.state.data.get('deps').map(function(dep){
        return <DepItem key={dep.get('name')} dep={dep}/> 
      })}

      <div className="ui bottom attached segment"></div>
    </div>;
  }
});