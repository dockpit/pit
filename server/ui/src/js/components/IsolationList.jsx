var React = require('react');
var Immutable = require('immutable')

var IsolationItem = require('./IsolationItem.jsx')

module.exports = React.createClass({
  render: function() {
  	var me = this  	
	var left = Immutable.List()
	var right = Immutable.List()
	this.props.isolations.forEach(function(iso, i){
		if(i & 1) {
		    right = right.push(iso)
		    return
		} 
		left = left.push(iso)
	})

    return <div className="ui two column grid doubling row">
		<div className="column">
			<div className="ui cards">
			{left.map(function(iso){			
				{if (me.props.selection === iso) {
					return <IsolationItem deps={me.props.deps} isSelected={true} key={iso.get('id')} isolation={iso}/> 
				} else{
					return <IsolationItem deps={me.props.deps} isSelected={false} key={iso.get('id')} isolation={iso}/> 
				}}
			})}
	    	</div>
		</div>
		<div className="column"> 
			<div className="ui cards">
			{right.map(function(iso){			
				{if (me.props.selection === iso) {
					return <IsolationItem deps={me.props.deps} isSelected={true} key={iso.get('id')} isolation={iso}/> 
				} else{
					return <IsolationItem deps={me.props.deps} isSelected={false} key={iso.get('id')} isolation={iso}/> 
				}}
			})}
	    	</div>
    	</div>
	</div>   
  }
});