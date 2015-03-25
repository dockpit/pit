var React = require('react');

var DepActions  = require('../actions/DepActions')
var DepStore = require('../stores/DepStore')
var DepList = require("./DepList.jsx")
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
    return <div>
      <h2>Dependencies<button className="dp-add" onClick={this.openDepForm}>+</button></h2>
      { this.state.showForm ? <DepForm closeFormFn={this.closeDepForm} dep={this.state.data.get('selection')}/> : null }

      <DepList deps={this.state.data.get('deps')}/>
    </div>;
  }
});