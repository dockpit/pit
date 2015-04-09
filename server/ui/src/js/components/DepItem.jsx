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
				if(dname === dep.get('id') && sname === st.get('id')) {
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
				<a className="header" href={"/deps/"+dep.get('id')+"/states/"+st.get('id')}>
					<h4>{st.get('name')}</h4>				
				</a>
				<div className="description">Image: {imageName ? <span>'(image: '+{imageName}+')'</span> : '<none>'}</div>			
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
		var conf = {
		    inline   : true,
		    on    : 'click',
		    position : 'right center',
		    target: $(React.findDOMNode(me.refs.addStateBtn)),
		    onVisible: function() {
		    	React.findDOMNode(me.refs.stateNameInput).focus()
		    }
		  }

		$(React.findDOMNode(this.refs.addStateArea)).popup(conf);
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
		$(React.findDOMNode(this.refs.addStateArea)).popup('hide')
	},

	render: function() {
		var me = this
		var dep = this.props.dep
		
		return <div onMouseEnter={this.enter} onMouseLeave={this.leave} style={{position: 'relative'}}>
			<h3 className="ui top attached header">
				{dep.get('name')} 				
			</h3>
			{this.state.hover ? <button onClick={this.removeDep} style={{position: 'absolute', top: '10px', right: '10px'}} className="circular ui compact basic icon small button"><i className="trash icon"></i></button> : null }
			<div className="ui list attached segment">
				{dep.get('states').map(function(st){
					return <DepStateItem isolations={me.props.isolations} key={dep.get('name')+st.get('name')} dep={dep} state={st}/>
				})}			
			</div>
			<div className="ui attached bottom secondary segment">
			  
				<div ref="addStateArea">
					<button ref="addStateBtn" className="basic tiny circular ui icon button">
					  <i className="icon plus"></i>
					</button>
					<div style={{width: '700px'}} className="ui popup">
						<form onSubmit={this.submitNewState} className="ui action input">
						  <input style={{border: '1px solid #DFDFDF'}} ref="stateNameInput" type="text" placeholder="e.g two users, no users"/>
						  <button type="submit" className="ui icon button">
						    <i className="save icon"></i>
						  </button>
						</form>
					</div>

					<a href="#" ref="addStateLink"> {'Add new ' + dep.get('name') + ' image...'}</a>
				</div>	

			</div>

		</div>
	}
});