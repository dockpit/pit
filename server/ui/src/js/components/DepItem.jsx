var React = require('react')

var DepActions  = require('../actions/DepActions')
var DepStateItem = React.createClass({
	removeDepState: function() {
		if(confirm("Are you sure you want to remove state '"+this.props.state.get('name')+"?")) {
			DepActions.removeDepState(this.props.dep, this.props.state)	
		}
	},

	render: function() {
		st = this.props.state
		dep = this.props.dep
		return <li>
			<a href={"/deps/"+dep.get('name')+"/states/"+st.get('name')}>{st.get('name')}</a>
			<button onClick={this.removeDepState}>x</button>
		</li>
	}
})

module.exports = React.createClass({
	removeDep: function() {
		if(confirm("Are you sure you want to remove dependency '"+this.props.dep.get('name')+"' and all its states?")) {
			DepActions.removeDep(this.props.dep)	
		}
	},

	render: function() {
		var dep = this.props.dep
		return <div className="ui attached segment" style={{position: 'relative'}}>		
			<div className="content">
				<h3 className="header">{dep.get('name')}</h3>
				<button style={{position: 'absolute', top: '10px', right: '10px'}} className="circular ui basic icon small button" onClick={this.removeDep}><i className="remove icon"></i></button>

				<ul className="dp-state-list">
					{dep.get('states').map(function(st){
						return <DepStateItem key={dep.get('name')+st.get('name')} dep={dep} state={st}/>
					})}
			    </ul>
				<a className="add-state-btn" href={"/deps/"+dep.get('name')+"/add-state"}>+ Add State</a>
			</div>
		</div>;
	}
});