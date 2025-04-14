// Use WebSocket transport endpoint.
const client = new Centrifuge('ws://localhost:8080/centrifugo/connection/websocket', {
   data: {
      username: "alex",
      password: "secret"
   }

});

client.connect()

const userInput = document.getElementById("username")
const chatInput = document.getElementById("chat")
const msgInput = document.getElementById("message")

const btnSubscribe = document.getElementById("subscribe")
const btnPublish = document.getElementById("publish")

const listView = document.getElementById("list")

btnSubscribe.addEventListener("click", (ev) =>
{
   if (userInput.value == "" || chatInput.value == "")
   {
      alert("Username and Chat must be not empty")

      return
   }

   const chatSub = client.newSubscription(chatInput.value)

   chatSub.on("publication", (ev) =>
   {
      let line = `<p><b>${ev.data.from}:</b> ${ev.data.message}</p>`

      listView.innerHTML += line
   })

   chatSub.subscribe()
})

btnPublish.addEventListener("click", (ev) =>
{
   if (msgInput.value == "")
   {
      return
   }

   client.publish(chatInput.value, {
      from: userInput.value,
      message: msgInput.value
   })

   msgInput.value = ""
})
