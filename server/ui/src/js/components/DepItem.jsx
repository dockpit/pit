var React = require('react')
var Immutable = require('immutable')

var DepActions  = require('../actions/DepActions')
var DepStateItem = React.createClass({
	getInitialState: function() {
		return {
			hover: false
		}
	},

	removeDepState: function() {
		if(confirm("Are you sure you want to remove state '"+this.props.state.get('name')+"?")) {
			DepActions.removeDepState(this.props.dep, this.props.state)	
		}
	},

	enter: function() { this.setState({hover: true})},
	leave: function() { this.setState({hover: false})},

	render: function() {
		st = this.props.state
		dep = this.props.dep
		imageName = st.get('image_name')

		var isolations = this.props.isolations.filter(function(iso){
			var res = false
			iso.get('states').forEach(function(sname, dname){
				if(dname === dep.get('name') && sname === st.get('name')) {
					res = true
					return
				}
			})
			return res
		})

		return <div className="item" onMouseEnter={this.enter} onMouseLeave={this.leave}>
			{this.state.hover ? <button style={{margin: '5px 5px 0 0'}} className="right floated circular compact ui red icon mini button" onClick={this.removeDepState}><i className="trash icon"></i></button> : null }
			<i className="angle right icon"></i>
			<div className="content">
				<a className="header" href={"/deps/"+dep.get('name')+"/states/"+st.get('id')}>
					{st.get('name')}
					{imageName ? <span> (image: {imageName})</span> : null}
				</a>
				<div className="description">In isolations: {isolations.size > 0 ? isolations.map(function(iso, i){
					return <em key={i}>{i ? ', ' : null}'{iso.get('name')}'</em>
				}) : <em>none</em>}</div>				
			</div>			
		</div>
	}
})

module.exports = React.createClass({
	getInitialState: function() {
		return {
			hover: false
		}
	},

	componentDidMount: function() {
		var me = this
		$(React.findDOMNode(this.refs.addStateBtn))
		  .popup({
		    inline   : true,
		    on    : 'click',
		    position : 'right center',
		    onVisible: function() {
		    	React.findDOMNode(me.refs.stateNameInput).focus()
		    }
		  })
		;
	},

	removeDep: function() {
		if(confirm("Are you sure you want to remove dependency '"+this.props.dep.get('name')+"' and all its states?")) {
			DepActions.removeDep(this.props.dep)	
		}
	},

	enter: function() { this.setState({hover: true})},
	leave: function() { this.setState({hover: false})},

	submitNewState: function(ev) {
		ev.preventDefault()
		var sname = React.findDOMNode(this.refs.stateNameInput).value
		if (sname) {
			var state = Immutable.Map({name: sname})

			DepActions.addDepState(this.props.dep, state)
		}

		React.findDOMNode(this.refs.stateNameInput).value = ''
		$(React.findDOMNode(this.refs.addStateBtn)).popup('hide')
	},

	render: function() {
		var me = this
		var dep = this.props.dep
		return <div className="ui attached segment" style={{position: 'relative'}}>		
			<div className="content">
				<h3 onMouseEnter={this.enter} onMouseLeave={this.leave} className="header">{dep.get('name')} {this.state.hover ? <button onClick={this.removeDep} style={{margin: '0px 0px -10px 10px'}} className="circular ui compact red icon small button"><i className="trash icon"></i></button> : null }</h3>
				<div className="ui list">
					{dep.get('states').map(function(st){
						return <DepStateItem isolations={me.props.isolations} key={dep.get('name')+st.get('name')} dep={dep} state={st}/>
					})}
					<div style={{paddingLeft: '20px'}} className="item dp-add-state">
						<a ref="addStateBtn"><i style={{fontSize: '0.65em'}} className="plus icon"></i>Add {me.props.dep.get('name')} state...</a>

						<div style={{width: '700px'}} className="ui popup">
							<form onSubmit={this.submitNewState} className="ui action input">
							  <input style={{border: '1px solid #DFDFDF'}} ref="stateNameInput" type="text" placeholder="e.g two users, no users"/>
							  <button type="submit" className="ui icon button">
							    <i className="save icon"></i>
							  </button>
							</form>
						</div>

					</div>
			    </div>
				
			</div>
		</div>;
	}
});