var React = require('react')
var Immutable = require('immutable')

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
			{ this.props.isSelected ? <td><button onClick={this.removeIsolationState} className="circular ui basic icon button mini"><i className="remove icon"></i></button></td> : null }
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
		return <div className="ui small input">
			<input style={{width: '50%', marginRight: '5px'}} ref="nameInput" defaultValue={this.props.isolation.get('name')} />
			<button  className="ui small compact button" onClick={this.commit}>Save</button> 
			or 
			<a style={{paddingLeft: '5px', fontWeight: 'normal'}} onClick={this.props.stopEditFn}>cancel</a>
		</div>
	}
})

var DependencySelector = React.createClass({
	componentDidMount: function() {
		var me = this
		$(React.findDOMNode(this.refs.dropDown)).dropdown({
			onChange: me.depSelected  
		})
	},

	depSelected: function(val) {
		var vals = val.split("::")
		IsolationActions.addState(this.props.isolation, vals[0], vals[1]) //: function(isolation, dname, sname)
		$(React.findDOMNode(this.refs.dropDown)).dropdown('restore defaults')
	},

	render: function() {
		var items = Immutable.List()
		this.props.deps.forEach(function(dep){
			items = items.push(Immutable.Map({classes: "header", text: dep.get('name'), key: dep.get('name')}))
			dep.get('states').forEach(function(st){
				items = items.push(Immutable.Map({classes: "item", value: dep.get('name')+'::'+st.get('name'), text: "  "+st.get('name'), key: dep.get('name')+st.get('name')}))
			})
		})

		return  <div ref="dropDown" style={{backgroundColor: "#CCC"}} className="ui floating dropdown bottom attached search selection button">
			<input type="hidden"/>
			<i className="plus icon"></i>
			<span className="text">Add Dependency</span>
	        <div className="menu">
	        {items.map(function(item){		
	        	return <div key={item.get('key')} data-value={item.get('value')} className={item.get('classes')}>
						{item.get('text')}
			    </div>
	        })}
	        </div>
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
		var classes = "isolation ui card fluid"
		if (this.props.isSelected) {
			classes = classes + " red"
		}

		var states = this.props.isolation.get('states').map(function(sname, dname){
			return {sname: sname, dname: dname}
		}).toSetSeq()

		return <div className={classes} onClick={this.selectIsolation}>
			<div className="content">
				{this.props.isSelected && !this.state.showNameEdit ? <button style={{position: 'absolute', top: '10px', right: '10px'}} onClick={this.removeIsolation} className="circular ui basic icon button mini"><i className="trash icon"></i></button> : null }
				{this.props.isSelected && !this.state.showNameEdit ? <button style={{position: 'absolute', top: '10px', right: '40px'}} onClick={this.startIsolationNameEdit} className="circular ui basic icon button mini"><i className="edit icon"></i></button> : null }
				<h3 className="header">
					{this.state.showNameEdit && this.props.isSelected ? <IsolationNameEditForm stopEditFn={this.stopIsolationNameEdit} isolation={this.props.isolation}/> : <span>{this.props.isolation.get('name')}</span> }	
				</h3>
				
				<div className="description">
					<table className="content ui very basic table">
						<tbody>
							{states.map(function(st){
								return 	<IsolationStateItem 
											isolation={me.props.isolation}
											isSelected={me.props.isSelected} 
											key={st.dname+st.sname} 
											state={st}/>	
							})}
						</tbody>
					</table>

					
				</div>
			</div>
			<DependencySelector isolation={this.props.isolation} deps={this.props.deps}/>
		</div>;
	}
});