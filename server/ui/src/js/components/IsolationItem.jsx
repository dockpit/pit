var React = require('react')

module.exports = React.createClass({
  render: function() {
  	return <li>{this.props.isolation.name}</li>;
  }
});