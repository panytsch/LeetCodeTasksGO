(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{35:function(e,t,n){e.exports=n(70)},40:function(e,t,n){},48:function(e,t,n){},70:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n.n(a),u=n(29),o=n.n(u),l=(n(40),n(4)),i=n(5),s=n(7),c=n(6),m=n(8),p=n(10),d=n(73),I=n(72),b=(n(46),n(48),n(12)),y=n(13),h=n(31),E=n(32);var f=Object(y.combineReducers)({userData:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{data:{type:null,name:null,userId:null},currentGame:{1:null,2:null,3:null,4:null,5:null,6:null,7:null,8:null,9:null},game:{myTurn:null,gameId:null,winner:null,pat:null,opponentName:null},timeoutId:null,isMyTurnTimeoutId:null,winTimeout:null},t=arguments.length>1?arguments[1]:void 0;switch(t.type){case"SET_NAME":return e.data.name=t.payload.name,e.data.userId=t.payload.userId,Object(b.a)({},e);case"FOUNDED_GAME":return e.data.type=t.payload.type,e.game.myTurn="x"===t.payload.type,e.game.gameId=t.payload.gameId,e.game.opponentName=t.payload.opponentName,e.game.myTurn||(e.isMyTurnTimeoutId=!0),clearInterval(e.timeoutId),e.timeoutId=null,Object(b.a)({},e);case"PENDING_GAME":return e.game.gameId=t.payload.gameId,e;case"SAVE_TIMEOUT_OTHER":return clearInterval(e.isMyTurnTimeoutId),e.isMyTurnTimeoutId=t.timeoutId,e;case"SAVE_TIMEOUT":return clearInterval(e.timeoutId),e.timeoutId=t.timeoutId,e;case"CLEAR_INTERVAL":return clearInterval(e.timeoutId),e.timeoutId=null,e;case"LEAVE_GAME":return e.data.type=null,e.game.gameId=null,e.game.myTurn=null,e.currentGame={1:null,2:null,3:null,4:null,5:null,6:null,7:null,8:null,9:null},e.game={myTurn:null,gameId:null,winner:null,pat:null,opponentName:null},e.isMyTurnTimeoutId&&clearInterval(e.isMyTurnTimeoutId),e.timeoutId&&clearInterval(e.timeoutId),e.isMyTurnTimeoutId=null,e.timeoutId=null,Object(b.a)({},e);case"MAKE_TURN":return e.currentGame[t.payload.itemNumber]=e.data.type,e.game.myTurn=!1,e.timeoutId=t.payload.timeoutId,Object(b.a)({},e);case"MY_TURN_FETCHED":return t.payload.me&&t.payload.me.map(function(t){e.currentGame[t]=e.data.type}),t.payload.opponent&&t.payload.opponent.map(function(t){e.currentGame[t]="o"===e.data.type?"x":"o"}),e.game.myTurn=!0,e.isMyTurnTimeoutId&&clearInterval(e.isMyTurnTimeoutId),e.timeoutId&&clearInterval(e.timeoutId),e.isMyTurnTimeoutId=null,Object(b.a)({},e);case"GAME_HAS_WINNER":return e.isMyTurnTimeoutId&&clearInterval(e.isMyTurnTimeoutId),e.timeoutId&&clearInterval(e.timeoutId),e.game.winner=t.payload.winner,Object(b.a)({},e);case"PAT_GAME":return e.isMyTurnTimeoutId&&clearInterval(e.isMyTurnTimeoutId),e.timeoutId&&clearInterval(e.timeoutId),e.game.pat=!0,Object(b.a)({},e);case"NEW_GAME":return e.timeoutId=t.payload.timeoutId,e.data.type=null,e.currentGame={1:null,2:null,3:null,4:null,5:null,6:null,7:null,8:null,9:null},e.game={myTurn:null,gameId:null,winner:null,pat:null,opponentName:null},e.isMyTurnTimeoutId=null,Object(b.a)({},e);default:return e}}}),T=Object(y.createStore)(f,Object(h.composeWithDevTools)(Object(y.applyMiddleware)(E.a))),g=n(74),v=function(e){function t(e){var n;return Object(l.a)(this,t),(n=Object(s.a)(this,Object(c.a)(t).call(this,e))).state={disabled:!1},n}return Object(m.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){return r.a.createElement("button",{className:this.props.classList&&this.props.classList.join(" ")||"btn btn-light",disabled:this.state.disabled,onClick:this.props.onClick||function(){}},this.props.text)}}]),t}(a.Component),O=n(71),j=function(e){function t(){return Object(l.a)(this,t),Object(s.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(m.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){return r.a.createElement(O.a,{to:this.props.route||"/",className:this.props.classList&&this.props.classList.join(" ")||"btn btn-light",onClick:this.props.onClick||function(){}},this.props.text)}}]),t}(a.Component),N=n(14),M=n.n(N),w="http://localhost:8000/api/",A={host:w,setName:function(e){return function(t){M.a.post("".concat(w,"users/new"),{name:e}).then(function(n){var a=n.data;a.status&&t({type:"SET_NAME",payload:{name:e,userId:a.userId}})})}},searchOpponent:function(e,t){return function(n){M.a.post("".concat(w,"games/join"),{userId:e,gameId:t}).then(function(e){var t=e.data;console.log(t),t.status&&t.type&&t.gameId?n({type:"FOUNDED_GAME",payload:{gameId:t.gameId,type:t.type,opponentName:t.opponentName}}):t.pending&&t.gameId&&n({type:"PENDING_GAME",payload:{gameId:t.gameId}})})}},leaveGame:function(e,t){return function(n){M.a.post("".concat(w,"games/leave"),{userId:e,gameId:t}).then(function(){n({type:"LEAVE_GAME"})})}},makeTurn:function(e,t,n,a){return function(r){M.a.post("".concat(w,"games/turn"),{userId:e,gameId:t,itemNumber:n}).then(function(e){var t=e.data;t.status&&(t.win?r({type:"GAME_HAS_WINNER",payload:{winner:t.winner}}):t.pat&&r({type:"PAT_GAME"}),r({type:"MAKE_TURN",payload:{itemNumber:n,timeoutId:a}}))})}},isMyTurn:function(e,t){return function(n){M.a.get("".concat(w,"games/request-to-turn"),{params:{userId:e,gameId:t}}).then(function(e){var t=e.data;t.status&&(t.win?n({type:"GAME_HAS_WINNER",payload:{winner:t.winner}}):t.pat&&n({type:"PAT_GAME"}),n({type:"MY_TURN_FETCHED",payload:{me:t.data.me,opponent:t.data.opponent}}))})}},isGameWin:function(e){return function(t){M.a.get("".concat(w,"games/request-to-win"),{params:{gameId:e}}).then(function(e){var n=e.data;n.status&&(n.win?t({type:"GAME_HAS_WINNER",payload:{winner:n.winner}}):n.pat&&t({type:"PAT_GAME"}))})}}},G=function(e){function t(e){return Object(l.a)(this,t),Object(s.a)(this,Object(c.a)(t).call(this,e))}return Object(m.a)(t,e),Object(i.a)(t,[{key:"componentDidMount",value:function(){this.props.clearInterval()}},{key:"render",value:function(){var e=this;return r.a.createElement("div",{className:"flex-column align-content-md-center"},r.a.createElement("div",{className:"flex-center-wrap",style:{width:"50%"}},r.a.createElement(v,{text:this.props.data.data.name||"Enter name",onClick:function(){return e.props.data.data.userId?function(e){e.preventDefault()}:e.props.setName(prompt("Your name"))},classList:["btn","btn-light","button-main"]}),r.a.createElement(j,{text:"New game",classList:["btn","btn-light","button-main"],route:"/new-game",onClick:function(){return e.props.data.data.userId?function(){}:e.props.setName(prompt("Your name"))}})))}}]),t}(r.a.Component),_=Object(p.b)(function(e){return{data:e.userData}},function(e){return{setName:function(t){return e(A.setName(t))},clearInterval:function(){return e({type:"CLEAR_INTERVAL"})}}})(Object(g.a)(G)),k={1:[0,1,1,0],2:[0,1,1,1],3:[0,0,1,1],4:[1,1,1,0],5:[1,1,1,1],6:[1,0,1,1],7:[1,1,0,0],8:[1,1,0,1],9:[1,0,0,1]},x=function(e){function t(e){return Object(l.a)(this,t),Object(s.a)(this,Object(c.a)(t).call(this,e))}return Object(m.a)(t,e),Object(i.a)(t,[{key:"myTurn",value:function(){var e=this,t=setInterval(function(){e.props.isMyTurn(e.props.data.data.userId,e.props.data.game.gameId)},1e3);this.props.makeTurn(this.props.data.data.userId,this.props.data.game.gameId,this.props.number,t)}},{key:"render",value:function(){var e=this,t=this.props.number;return r.a.createElement("div",{className:"game-board-item",style:{borderTop:k[t][0]&&"2px solid black",borderRight:k[t][1]&&"2px solid black",borderBottom:k[t][2]&&"2px solid black",borderLeft:k[t][3]&&"2px solid black"},onClick:this.props.data.game.myTurn&&!this.props.data.currentGame[t]?function(){return e.myTurn()}:function(){}},this.props.data.currentGame[t])}}]),t}(r.a.Component),C=Object(p.b)(function(e){return{data:e.userData}},function(e){return{makeTurn:function(t,n,a,r){return e(A.makeTurn(t,n,a,r))},isMyTurn:function(t,n){return e(A.isMyTurn(t,n))}}})(Object(g.a)(x)),L=function(e){function t(e){var n;Object(l.a)(this,t),n=Object(s.a)(this,Object(c.a)(t).call(this,e));var a=setInterval(function(){e.searchOpponent(e.data.data.userId,e.data.game.gameId||null)},2e3);return e.setTimeoutId(a),n}return Object(m.a)(t,e),Object(i.a)(t,[{key:"componentDidUpdate",value:function(){var e=this;if(setTimeout(function(){e.props.isGameWin(e.props.data.game.gameId)},7e4),!this.props.data.game.myTurn&&!0===this.props.data.isMyTurnTimeoutId){var t=setInterval(function(){e.props.isMyTurn(e.props.data.data.userId,e.props.data.game.gameId)},1e3);this.props.setTimeoutTurnId(t)}}},{key:"render",value:function(){var e=this,t=this.props.data.game,n=t.pat,a=t.winner,u=t.opponentName;return r.a.createElement("div",null,this.props.data.data.name&&u&&!n&&!a&&r.a.createElement(v,{text:"".concat(this.props.data.data.name," vs ").concat(this.props.data.game.opponentName),classList:["btn","btn-light","button-main"]}),a&&r.a.createElement(v,{text:"".concat(a," WIN!!!"),classList:["btn","btn-light","button-main"]}),n&&r.a.createElement(v,{text:"The game is a draw",classList:["btn","btn-light","button-main"]}),this.props.data.data.type&&r.a.createElement("h5",null,"Your letter is ",this.props.data.data.type),!n&&!a&&this.props.data.game.opponentName&&(this.props.data.game.myTurn?r.a.createElement("p",null,"Your turn"):r.a.createElement("p",null,"Opponent`s turn")),(n||a)&&r.a.createElement(v,{text:"New Game",onClick:function(){var t=setInterval(function(){e.props.searchOpponent(e.props.data.data.userId,e.props.data.game.gameId||null)},2e3);e.props.newGame(t)},classList:["btn","btn-light","button-main"]}),r.a.createElement("div",{className:"game-board-wrap"},r.a.createElement("div",{className:"point"},r.a.createElement(C,{number:1}),r.a.createElement(C,{number:4}),r.a.createElement(C,{number:7})),r.a.createElement("div",{className:"point"},r.a.createElement(C,{number:2}),r.a.createElement(C,{number:5}),r.a.createElement(C,{number:8})),r.a.createElement("div",{className:"point"},r.a.createElement(C,{number:3}),r.a.createElement(C,{number:6}),r.a.createElement(C,{number:9}))))}}]),t}(r.a.Component),D=Object(p.b)(function(e){return{data:e.userData}},function(e){return{searchOpponent:function(t,n){return e(A.searchOpponent(t,n))},setTimeoutId:function(t){return e({type:"SAVE_TIMEOUT",timeoutId:t})},setTimeoutTurnId:function(t){return e({type:"SAVE_TIMEOUT_OTHER",timeoutId:t})},isMyTurn:function(t,n){return e(A.isMyTurn(t,n))},newGame:function(t){return e({type:"NEW_GAME",payload:{timeoutId:t}})},isGameWin:function(t){return e(A.isGameWin(t))}}})(Object(g.a)(L)),R=function(e){function t(e){var n;return Object(l.a)(this,t),(n=Object(s.a)(this,Object(c.a)(t).call(this,e))).props.data.data.userId||n.props.history.push("/"),n}return Object(m.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){var e=this;return r.a.createElement("div",{className:"flex-column align-content-md-center"},r.a.createElement("div",{className:"flex-center-wrap",style:{width:"50%"}},r.a.createElement(D,null),r.a.createElement(j,{text:"Abandon game",classList:["btn","btn-light","button-main"],route:"/",onClick:function(){return e.props.leaveGame(e.props.data.data.userId,e.props.data.game.gameId)}})))}}]),t}(r.a.Component),W=Object(p.b)(function(e){return{data:e.userData}},function(e){return{leaveGame:function(t,n){return e(A.leaveGame(t,n))}}})(Object(g.a)(R)),S=function(e){function t(){return Object(l.a)(this,t),Object(s.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(m.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){return r.a.createElement(p.a,{store:T},r.a.createElement(d.a,null,r.a.createElement("div",null,r.a.createElement(I.a,{exact:!0,path:"/",component:_}),r.a.createElement(I.a,{path:"/new-game",component:W}))))}}]),t}(a.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(r.a.createElement(S,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}},[[35,2,1]]]);
//# sourceMappingURL=main.75d40faa.chunk.js.map