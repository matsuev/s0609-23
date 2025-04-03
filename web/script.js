// Use WebSocket transport endpoint.
const client = new Centrifuge('ws://loalhost:8080/centrifugo/connection/websocket');

// Allocate Subscription to a channel.
const newsSub = client.newSubscription('news');

const list = document.getElementById("list")

// React on `news` channel real-time publications.
newsSub.on('publication', function (ctx)
{
   console.log(ctx.data);

   list.innerHTML += "<p>" + ctx.data.message + "</p>"
});

// Trigger subscribe process.
newsSub.subscribe();

// Trigger actual connection establishement.
client.connect();