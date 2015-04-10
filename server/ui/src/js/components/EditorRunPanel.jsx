var React = require('react');
var Immutable = require('immutable')

var EditorActions  = require('../actions/EditorActions')
var EditorStore = require('../stores/EditorStore')

var EditorRunContainerInfo = React.createClass({
	render: function() {
		var me = this
		var noports = <div className="ui horizontal divider">no exposed ports</div>
		var bindings = []

		if (!this.props.info.get('NetworkSettings').get("Ports") ){
			return noports
		}

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
			return noports
		}

		return <div>
			<div className="ui horizontal divider">exposed ports:</div>

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
			    {this.props.build.get('is_running') ? <div style={{fontSize: '1em', position: "absolute", top: "10px", left: "10px"}} className="ui active icon inline loader"></div> : <img style={{position: "absolute", top: 0, left: "-4px", opacity: (this.props.build.get('is_running') ? 0.3 : 1)}} src="/static/src/img/icon_build_image@1x.png"/>} 

			    <div style={{paddingLeft: '30px'}} className="content">
			      <div className="title">1. Build</div>
			      <div className="description">Create a Docker Image</div>
			    </div>
			  </a>
			  <a onClick={this.startRun} className={'step' + (!this.props.build.get('image_name') || runRunning  ? ' disabled' : '')}>
			  	{runRunning ? <div style={{fontSize: '1em', position: "absolute", top: "10px", left: "20px"}} className="ui active icon inline loader"></div> : <img style={{position: "absolute", top: "3px", left: "10px", zIndex: 100, opacity: (!this.props.build.get('image_name') || runRunning ? 0.3 : 1)}} src="/static/src/img/icon_run_image@1x.png"/>} 
			
			    <div style={{paddingLeft: '30px'}} className="content">
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

			<div onClick={this.props.completedFn} className={'ui fluid icon primary button'+ (this.props.run.get('is_ready') ? '' : ' disabled')}>
				3. Complete Image <i className="right checkmark icon"></i>
			</div>      	
		</div>
}})

module.exports = EditorRunPanel