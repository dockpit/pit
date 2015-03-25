var React = require('react')
var Immutable = require('immutable')

var IsolationActions  = require('../actions/IsolationActions')

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

		var newiso = Immutable.Map({name: name})
		IsolationActions.create(newiso)
		React.findDOMNode(this.refs.nameInput).value = ''
		this.props.closeFormFn(ev)
	},

	render: function(){
		return <form onSubmit={this.handleSubmit}>
			<input ref="nameInput" type="text" id="isolation_name" name="name" required/>			
			<button type="submit">Create isolation</button>
			<button onClick={this.props.closeFormFn}>Cancel</button>		
		</form>;
	}
})