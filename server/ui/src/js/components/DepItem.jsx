var React = require('react')

module.exports = React.createClass({
  render: function() {
  	return <li>
  		{this.props.dep.get('name')}
  		<a href={"/deps/"+this.props.dep.get('name')+"/add-state"}>+ Add State</a>
  	</li>;
  }
});