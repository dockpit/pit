var React = require('react')

module.exports = React.createClass({
	render: function(){
		return <form action="/isolations" enctype="application/x-www-form-urlencoded" method="post">
			<input type="text" id="isolation_name" name="name"/>			
			<button onClick={this.props.closeFormFn} type="submit">Cancel</button>
			<button type="submit">Create isolation</button>
		</form>;
	}
})