var React = require('react')

module.exports = React.createClass({
	componentDidMount: function() {
		React.findDOMNode(this.refs.nameInput).focus();  
	},

	render: function(){
		return <form action="/deps" encType="application/x-www-form-urlencoded" method="post">
			<input ref="nameInput" type="text" id="dep_name" name="name"/>				
			<button type="submit">Create dep</button>
			<button onClick={this.props.closeFormFn}>Cancel</button>
		</form>;
	}
})