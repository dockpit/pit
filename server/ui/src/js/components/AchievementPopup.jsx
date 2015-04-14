var React = require('react')
var Immutable = require('immutable')

var StatActions  = require('../actions/StatActions')

module.exports = React.createClass({
	getInitialState: function() {
		return {
			frame: 0,
		}
	},

	preventClickPropagation: function(ev){
		ev.stopPropagation()
	},

	componentDidMount: function() {
		var me = this
		this.interval = setInterval(function(){
			var curr = me.state.frame + 1
			if (curr == 4) {
				curr = 0
			}

			me.setState({frame: curr})
		}, 2000)
	},

	componentDidUnmount: function() {
		clearInterval(this.interval)
	},

	close: function(achievement) {
		StatActions.shown(achievement)
	},

	render: function(){
		var me = this
		var show = Immutable.List()
		this.props.achievements.forEach(function(a){
			if(a.get('shown') === true) { return }
			show = show.push(a)
		})

		if (show.size == 0) {
			return null
		}

		//@todo create modals for various different achievements
		var modals = show.map(function(achievement, i) {
			return <div key={i} onClick={me.close.bind(me, achievement)} className="ui dimmer modals page transition visible active">
				<div onClick={me.preventClickPropagation} style={{top: '150px'}} className="ui modal transition small visible active">

					<h2 className="ui header">
						<img src="/static/src/img/icon_dep@1x.png"/>
						<div className="content">
						  Nicely done!
						</div>
					</h2>


				    <div className="content">
				    	{"You've just setup the first dependency of the default isolation, next you could:"}
				    	<ol>
				    		<li>
				    			{"Commit you isolation for other team members to use:"}
				    			<pre style={{backgroundColor: "#333", color: "#CCC", padding: "10px"}}>
				    				{"$> git add dockpit.db\n"}
				    				{"$> git commit -m \"initial isolation\""}
				    			</pre>
				    		</li>
				    		<li>
				    			{"Start your isolation by heading over to your terminal UI:"}
				    			<pre style={{backgroundColor: "#333", color: "#CCC", padding: "10px"}}>
{"Docker Host: OK (tcp://192.168.99.100:2376)\n"}
{"Web Interface: Ok (http://[::]:3838)\n"}
{"Current Isolation: "+(me.state.frame == 0 ? '' : 'default')+" ("+(me.state.frame == 0 ? '<none>' : (me.state.frame == 1 ? 'starting...' : 'OK'))+")\n\n"}
{"Isolations:\n\n"}
{"   ["+(me.state.frame < 3 ? '*' : ' ')+"] default\n"}
{"   ["+(me.state.frame > 2 ? '*' : ' ')+"] Stop all isolations\n\n"}
{"Use Esc or Ctrl-C to exit"}
				    			</pre>
				    		</li>
				    	</ol>
					</div>

					<div className="actions">						
						<button className="ui button icon primary" onClick={me.close.bind(me, achievement)}><i className="icon checkmark"></i> Done</button>		
					</div>
				</div>
			</div>
		})

		return <div>{modals}</div>
	}
})