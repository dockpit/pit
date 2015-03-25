var React = require('react')
var Immutable = require('immutable')

var DepActions  = require('../actions/DepActions')

module.exports = React.createClass({
	componentDidMount: function() {
		React.findDOMNode(this.refs.nameInput).focus();  
	},

	handleSubmit: function(ev) {
		ev.preventDefault();
		name = React.findDOMNode(this.refs.nameInput).value
		if (name == "") {
			return
		}

		var newdep = Immutable.Map({name: name})
		DepActions.create(newdep)
		React.findDOMNode(this.refs.nameInput).value = ''
		this.props.closeFormFn(ev)
	},	

	render: function(){
		return <form onSubmit={this.handleSubmit}>
			<input ref="nameInput" type="text" id="dep_name" name="name" required="required"/>				
			<button type="submit">Create dep</button>
			<button onClick={this.props.closeFormFn}>Cancel</button>
		</form>;
	}
})