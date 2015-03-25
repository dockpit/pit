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
		return <li>
			<div className="hr-text">
			  <h3>
			    {dep.get('name')}		    
			  </h3>
			  <button onClick={this.removeDep}>x</button>
			</div>
			<ul className="dp-state-list">
				{dep.get('states').map(function(st){
					return <DepStateItem key={dep.get('name')+st.get('name')} dep={dep} state={st}/>
				})}
		    </ul>
			<a className="add-state-btn" href={"/deps/"+dep.get('name')+"/add-state"}>+ Add State</a>
		</li>;
	}
});