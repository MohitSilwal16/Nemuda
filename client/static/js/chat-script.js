function clearMessageField() {
    document.getElementById("message").value = "";
}

var socket = new WebSocket("ws://localhost:3000/ws/chat");

socket.addEventListener("open", function (event) {
    console.log("WebSocket Connected");
});

socket.addEventListener("close", function (event) {
    console.log("WebSocket Disconnected");
    try {
        openAlert("WebSocket Disconnected", "Try Refreshing Browser");
    } catch (e) { }
});

// Listen for messages
socket.addEventListener("message", function (event) {
    const responseData = JSON.parse(event.data);
    const receiver = document.getElementById("title").innerText;

    if (receiver.error) {
        if (receiver.error == "Invalid Session Token") {
            location.href = "login";
        }
        openAlert("Error", receiver.error);
        return;
    }

    // There's no err message so continue
    // Send message in alert pop up if receiver hasn't open sender's chat
    if (receiver != responseData.sender && !responseData.selfMessage) {
        openNotification(responseData.sender, responseData.messageContent);
        return;
    }

    const messagesContainer = document.getElementById("messages");
    const messageCard = createMessageCard(responseData);
    messagesContainer.appendChild(messageCard);

    let noMsgBox = document.getElementById("noMsgBox");

    if (noMsgBox) {
        noMsgBox.remove();
        noMsgBox.classList.add("hidden");
    }

    // Scroll to the bottom after adding the new message
    scrollToBottom(messagesContainer);
});

// Send a message to the server
function onMessageSend() {
    event.preventDefault();
    const messageInput = document.getElementById("messageInput");
    const message = messageInput.value.trim();

    if (message == "") {
        openAlert("Empty Message", "Message is Empty");
        return;
    }

    const jsonData = {
        message: message,
        receiver: document.getElementById("title").innerText,
    };

    socket.send(JSON.stringify(jsonData));
    messageInput.value = "";
}

function createMessageCard(responseData) {
    // Create the main container div
    const messageContainer = document.createElement("div");
    messageContainer.classList.add("flex", "mb-4");

    // Create the inner message bubble div
    const messageBubble = document.createElement("div");
    messageBubble.classList.add("p-3", "rounded-lg", "max-w-xs");

    // Create the date-time element
    const dateTimeElement = document.createElement("div");
    dateTimeElement.classList.add("flex", "justify-between", "items-center");
    const dateTimeSpan = document.createElement("span");
    dateTimeSpan.classList.add("text-xs");
    dateTimeSpan.textContent = responseData.dateTime;
    dateTimeElement.appendChild(dateTimeSpan);

    // Create the message content element
    const messageContent = document.createElement("p");
    messageContent.classList.add("mt-1");
    messageContent.textContent = responseData.messageContent;

    // Create the status element
    const statusElement = document.createElement("div");
    statusElement.classList.add("mt-1", "text-right", "text-xs");
    statusElement.textContent = responseData.status;

    // Append elements to message bubble
    messageBubble.appendChild(dateTimeElement);
    messageBubble.appendChild(messageContent);
    messageBubble.appendChild(statusElement);

    messageBubble.classList.add("bg-blue-600", "text-white");

    // Conditionally style based on sender or receiver
    if (responseData.selfMessage) {
        messageContainer.classList.add("justify-end");
        statusElement.classList.add("text-gray-300");
    } else {
        messageContainer.classList.add("justify-start");
        statusElement.classList.add("hidden");
    }

    // Append the message bubble to the main container
    messageContainer.appendChild(messageBubble);

    return messageContainer;
}

function openAlert(title, body) {
    document.getElementById("overlay").style.display = "flex";
    document.getElementById("alertTitle").innerHTML = title;
    document.getElementById("alertBody").innerHTML = body;
}

function closeAlert() {
    document.getElementById("overlay").style.display = "none";
}

function openNotification(title, body) {
    document.getElementById("notificationTitle").innerText = title;
    document.getElementById("notificationBody").innerText = body;

    let notification = document.getElementById("notification");
    notification.style.display = "flex";
}

function onclickNotification() {
    const title = document.getElementById("notificationTitle").innerText;

    fetch("/message/" + title)
        .then((response) => response.text())
        .then((html) => {
            document.getElementById("message-body").innerHTML = html;
            closeNotification();
        })
        .catch((error) => {
            console.error("Error fetching data:", error);
        });

    const msgInput = document.getElementById("messageInput");
    msgInput.removeAttribute("readonly");
    msgInput.focus();
}

function closeNotification() {
    document.getElementById("notification").style.display = "none";
}

// Called when username card is clicked
function onUserClick() {
    messageInput.focus();
    document.getElementById("messageInput").removeAttribute("readonly");
}

// Close socket when back button is pressed
window.addEventListener('popstate', function (event) {
    socket.close();
});