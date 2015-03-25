var React = require('react')

var IsolationActions  = require('../actions/IsolationActions')

var IsolationStateItem = React.createClass({
	removeIsolationState: function() {
		IsolationActions.removeState(this.props.isolation, this.props.state)
	},

	render: function() {
		st = this.props.state
		return <tr>
			<td>{st.dname}:</td>
			<td>'{st.sname}'</td>
			{ this.props.isSelected ? <td><button onClick={this.removeIsolationState}>x</button></td> : null }
		</tr>	
	}
})

var IsolationNameEditForm = React.createClass({
	componentDidMount: function() {
		React.findDOMNode(this.refs.nameInput).focus();  
	},

	commit: function() {
		var val = React.findDOMNode(this.refs.nameInput).value
		if (val !== this.props.isolation.get('name')) {
			var newiso = this.props.isolation.set('name', val) 
			IsolationActions.updateName(newiso, this.props.isolation)					
		}

		this.props.stopEditFn()
	},

	render: function() {
		st = this.props.state
		return <div>
			<input ref="nameInput" defaultValue={this.props.isolation.get('name')} />
			<button onClick={this.commit}>commit</button> or <button onClick={this.props.stopEditFn}>cancel</button>
		</div>
	}
})

module.exports = React.createClass({
	getInitialState: function() {
		return {
			showNameEdit: false
		}
	},

	removeIsolation: function(ev) {
		ev.stopPropagation()
		if(confirm("Are you sure you want to remove isolation '"+this.props.isolation.get('name')+"'?")) {
			IsolationActions.remove(this.props.isolation)
		}
	},

	selectIsolation: function() {
		IsolationActions.select(this.props.isolation)
	},

	startIsolationNameEdit: function() {
		this.setState({showNameEdit: true})
	},

	stopIsolationNameEdit: function() {
		this.setState({showNameEdit: false})
	},

	render: function() {
		var me = this
		var classes = "isolation"
		if (this.props.isSelected) {
			classes = classes + " selected"
		}

		var states = this.props.isolation.get('states').map(function(sname, dname){
			return {sname: sname, dname: dname}
		}).toSetSeq()

		return <li className={classes} onClick={this.selectIsolation}>
			<h3>			
				{this.state.showNameEdit && this.props.isSelected ? <IsolationNameEditForm stopEditFn={this.stopIsolationNameEdit} isolation={this.props.isolation}/> : <span>{this.props.isolation.get('name')}</span> }
				{this.props.isSelected && !this.state.showNameEdit ? <button onClick={this.startIsolationNameEdit}>edit</button> : null }
				{this.props.isSelected && !this.state.showNameEdit ? <button onClick={this.removeIsolation}>x</button> : null }
			</h3>
			<table>
				<tbody>
					<tr><td rowSpan={states.size+1}>Given</td></tr>
					{states.map(function(st){
						return 	<IsolationStateItem 
									isolation={me.props.isolation}
									isSelected={me.props.isSelected} 
									key={st.dname+st.sname} 
									state={st}/>	
					})}
				</tbody>
			</table>

			{ this.props.isSelected ? <a href={"/isolations/"+this.props.isolation.get('id')+"/add-dep"}>+ Add Dependency State</a> : null }
		</li>;
	}
});