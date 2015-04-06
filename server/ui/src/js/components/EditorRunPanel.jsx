var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

var EditorRunContainerInfo = React.createClass({
	render: function() {
		var me = this

		var bindings = []
		this.props.info.get('NetworkSettings').get("Ports").forEach(function(hconfs, cport){
			if(!hconfs) {
				return
			} 

			bindings.push({
				hconfs: hconfs,
				cport: cport,
			})
		})

		if (bindings.length == 0) {
			return <div className="ui horizontal divider">no exposed ports</div>
		}

		return <div>
			<div className="ui horizontal divider">exposed:</div>

			<table className="ui very basic two column table">
			  <tbody>
				{bindings.map(function(conf, i) {
					return <tr key={i}>
				      <td className="center aligned">Container port {conf.cport}</td>				     
				      <td className="center aligned">Bound to {conf.hconfs.map(function(hconf, i){
				      	return <a key={i} target="_blank" href={'http://'+ me.props.dockerHostIp+':'+hconf.get('HostPort')}>{hconf.get('HostIp')}:{hconf.get('HostPort')}</a>
				      })}</td>
				    </tr>
				})}
			  </tbody>
			</table>

			<div className="ui divider"></div>
		</div>
	}
})

var EditorRunPanel = React.createClass({

	startBuild: function() {
		EditorActions.startBuild(this.props.depId, this.props.state)
	},

	startRun: function() {
		EditorActions.startRun(this.props.depId, this.props.state)
	},

	render: function(){
		var output = this.props.output		
		if(output) {
			output = window.atob(output)
		}		

		var runRunning = (this.props.run.get('id') && !this.props.run.get('is_ready') && !this.props.run.get('error'))

		return <div className="column">

			<div className="ui two steps attached top">
			  <a onClick={this.startBuild} className={'step' + (this.props.build.get('is_running') ? ' disabled': '')}>
			    {this.props.build.get('is_running') ? <div style={{fontSize: '1em'}} className="ui active icon inline loader"></div> : <i className="repeat icon"></i>} 

			    <div className="content">
			      <div className="title">1. Build</div>
			      <div className="description">Create a Docker Image</div>
			    </div>
			  </a>
			  <a onClick={this.startRun} className={'step' + (!this.props.build.get('image_name') || runRunning  ? ' disabled' : '')}>
			  	{runRunning ? <div style={{fontSize: '1em'}} className="ui active icon inline loader"></div> : <i className="terminal icon"></i>} 
			
			    <div className="content">
			      <div className="title">2. Test</div>
			      <div className="description">Run the Image</div>
			    </div>
			  </a>
			</div>

			<div className="ui attached bottom segment">			
				{this.props.error ? <div className="ui negative message"><p>{this.props.error}</p></div> : null }

				<pre>
					{output}
				</pre>
			</div>		


			{this.props.run.get('container_info') ? <EditorRunContainerInfo dockerHostIp={this.props.dockerHostIp} info={this.props.run.get('container_info')}/> : <div className="ui divider"></div>}


			<div onClick={this.props.completedFn} className={'ui fluid icon positive button'+ (this.props.run.get('is_ready') ? '' : ' disabled')}>
				3. Complete State <i className="right checkmark icon"></i>
			</div>      	
		</div>
}})

module.exports = EditorRunPanel