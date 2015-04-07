var React = require('react')
var Immutable = require('immutable')

var StatActions  = require('../actions/StatActions')

module.exports = React.createClass({
	preventClickPropagation: function(ev){
		ev.stopPropagation()
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
					<div className="header">
				     	Nicely done!
				    </div>

				    <div className="content">
				    	{"You've just setup the first dependency of your default fixture, next you could:"}
				    	<ol>
				    		<li>{"Commit the fixture for other team members"}</li>
				    		<li>{"Start your fixture by..."}</li>
				    	</ol>
					</div>

					<div className="actions">						
						<button className="ui button icon positive" onClick={me.close.bind(me, achievement)}><i className="icon checkmark"></i> Done</button>		
					</div>
				</div>
			</div>
		})

		return <div>{modals}</div>
	}
})