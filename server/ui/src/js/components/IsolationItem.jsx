var React = require('react')

module.exports = React.createClass({
  render: function() {
  	return <li>
  		{this.props.isolation.name}
  		<a href={"/isolations/"+this.props.isolation.name+"/add-dep"}>+ Add dep</a>
  	</li>;
  }
});