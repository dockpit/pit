var React = require('react')

var IsolationActions  = require('../actions/IsolationActions')

module.exports = React.createClass({
	selectIsolation: function() {
		IsolationActions.select(this.props.isolation)
	},

	removeIsolation: function() {
		if(confirm("Are you sure?")) {
			IsolationActions.remove(this.props.isolation)	
		}
	},

	render: function() {
		var classes = "isolation"
		if (this.props.isSelected) {
			classes = classes + " selected"
		}

		return <li className={classes} onClick={this.selectIsolation}>
			{this.props.isolation.get('name')}
			<a href={"/isolations/"+this.props.isolation.get('name')+"/add-dep"}>+ Add dep</a>
			{ this.props.isSelected ? <button onClick={this.removeIsolation}>Remove</button> : null }
		</li>;
	}
});