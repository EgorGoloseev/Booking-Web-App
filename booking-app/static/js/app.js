async function loadRooms(){

    const res = await fetch("/rooms")

    const rooms = await res.json()

    const container = document.getElementById("rooms")

    container.innerHTML=""

    rooms.forEach(room=>{

        container.innerHTML += `

<div class="room">

<h3>${room.name}</h3>

<p>📍 ${room.location}</p>

<p>👥 Вместимость: ${room.capacity}</p>

<button onclick="bookRoom(${room.id})">Забронировать</button>

</div>

`

    })

}

function bookRoom(id){

    alert("Форма бронирования будет здесь. Комната: "+id)

}