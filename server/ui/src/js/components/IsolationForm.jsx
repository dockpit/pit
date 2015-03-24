var React = require('react')

module.exports = React.createClass({

	componentDidMount: function() {
		React.findDOMNode(this.refs.nameInput).focus();  
	},

	render: function(){
		return <form action="/isolations" encType="application/x-www-form-urlencoded" method="post">
			<input ref="nameInput" type="text" id="isolation_name" name="name"/>			
			<button type="submit">Create isolation</button>
			<button onClick={this.props.closeFormFn}>Cancel</button>		
		</form>;
	}
})