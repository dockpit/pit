var React = require('react')

var IsolationActions  = require('../actions/IsolationActions')
var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
	selectIsolation: function() {
		window.location = "/isolations/"+this.props.isolation.get('name')
	},

	render: function() {
		var classes = "isolation"
		if (this.props.isSelected) {
			classes = classes + " selected"
		}

		var states = this.props.isolation.get('states').map(function(sname, dname){
			return {sname: sname, dname: dname}
		}).toSetSeq()

		return <li className={classes} onClick={this.selectIsolation}>
			<h3>{this.props.isolation.get('name')}</h3>
			<hr />

			<table>
				<tr><td rowSpan={states.size+1}>Given</td></tr>
				{states.map(function(st){
					return <tr key={st.dname+st.sname}>
						<td>{st.dname}:</td>
						<td>'{st.sname}'</td>
					</tr>				
				})}
			</table>
		</li>;
	}
});