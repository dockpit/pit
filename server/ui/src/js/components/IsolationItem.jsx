var React = require('react')
var Immutable = require('immutable')

var IsolationActions  = require('../actions/IsolationActions')

var IsolationStateItem = React.createClass({
	removeIsolationState: function() {
		IsolationActions.removeState(this.props.isolation, this.props.state)
	},

	render: function() {
		st = this.props.state
		this.props.deps.forEach(function(dep){			
			if(st.depid == dep.get('id')) {				
				dep.get('states').forEach(function(sst){
					if (st.stateid == sst.get('id')) {
						st.state = sst
						st.dep = dep						
						return
					}
				})
				return
			}
		})

		//only render if we got dep info
		if (!st.dep) {
			return <tr></tr>
		}

		return <tr>
			<td width="20"><img style={{opacity: 0.2}} height="20" src={depIconPath(st.dep, "0,7")} /></td>
			<td>{st.dep.get('name')} @ <em>{st.state.get('name')}</em></td>
			{ this.props.isSelected ? <td width="50" style={{textAlign: 'right'}}><button onClick={this.removeIsolationState} className="circular ui basic icon button mini"><i className="remove icon"></i></button></td> : null }
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
		IsolationActions.addState(this.props.isolation, vals[0], vals[1])
		$(React.findDOMNode(this.refs.dropDown)).dropdown('restore defaults')
	},

	render: function() {
		var me = this
		var items = Immutable.List()
		this.props.deps.forEach(function(dep){
			var alreadyHasDep = false
			me.props.isolation.get('states').forEach(function(stateid, depid){
				if (depid == dep.get('id')) {
					alreadyHasDep = true
					return
				}
			})

			//only allow selection of a dep if isolation not longer has one
			if (alreadyHasDep) {
				return
			}

			items = items.push(Immutable.Map({
				dep: dep,
				classes: 	"header", 
				text: 		dep.get('name'), 
				key: 		dep.get('name')
			}))
			
			dep.get('states').forEach(function(st){
				items = items.push(Immutable.Map({
					classes: 	"item", 
					value: 		dep.get('id')+'::'+st.get('id'), 
					text: 		"  "+st.get('name'), 
					key: 		dep.get('id')+st.get('id')
				}))
			})
		})

		var mappeditems = items.map(function(item){		
        	return <div key={item.get('key')} data-value={item.get('value')} className={item.get('classes')}>
        			{item.get('dep') ? <img className="ui mini avatar image" src={depIconPath(item.get('dep'), "0,7")}/> : '- ' }
					{item.get('text')}
		    </div>
	    })

	    if(items.size == 0) {
	    	mappeditems = <div className="item"><em>All dependencies are represented already</em></div>
	    }

		return  <div ref="dropDown" className="ui inline labeled icon dropdown basic tiny compact button">
			<input type="hidden"/>
			<i className="plus icon"></i>
			<span style={{fontWeight: "normal", color: "#333"}} className="text">Add Image</span>
	        <div className="menu">
	        	{mappeditems}
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
		var states = this.props.isolation.get('states').map(function(stateid, depid){
			return {stateid: stateid, depid: depid}
		}).toSetSeq()

		var stcontent = <div style={{padding: '0px 10px'}} className={"content" + (this.props.isSelected ? " active" : "")}>
			<table style={{tableLayout: "fixed"}} className="ui very basic table">
				<tbody>
					{states.map(function(st){
						return 	<IsolationStateItem 
							isolation={me.props.isolation}
							isSelected={me.props.isSelected} 
							key={st.depid+st.stateid} 
							deps={me.props.deps}
							state={st}/>	
					})}
				</tbody>
				<tfoot>
					<tr><td></td><td><DependencySelector isolation={this.props.isolation} deps={this.props.deps}/></td></tr>
				</tfoot>
		    </table>
	    </div>

	    if (states.size == 0) {
	    	stcontent = <div className={"content" + (this.props.isSelected ? " active" : "")}>
	    	<span>{"Isolations may contain one specific image for each dependency to form a certain scenario: "}</span>
	    	<DependencySelector isolation={this.props.isolation} deps={this.props.deps}/>
	    </div>

	    }


		return <div style={{position: "relative"}} onClick={this.selectIsolation}>
		    <div className={"title" + (this.props.isSelected ? " active" : "")}>				
			  {this.props.isSelected && !this.state.showNameEdit && this.props.isolation.get('id') !== 'default' ? <button style={{position: 'absolute', top: '10px', right: '10px'}} onClick={this.removeIsolation} className="circular ui basic icon button mini"><i className="trash icon"></i></button> : null }
			  {this.props.isSelected && !this.state.showNameEdit && this.props.isolation.get('id') !== 'default' ? <button style={{position: 'absolute', top: '10px', right: '40px'}} onClick={this.startIsolationNameEdit} className="circular ui basic icon button mini"><i className="edit icon"></i></button> : null }

		      <i className="dropdown icon"></i>
		      {this.state.showNameEdit && this.props.isSelected ? <IsolationNameEditForm stopEditFn={this.stopIsolationNameEdit} isolation={this.props.isolation}/> : <span>{this.props.isolation.get('name')}</span> }	
		    </div>

		    {stcontent}

		    
		</div>
	}
});