async function loadRooms() {
    const res = await fetch("/rooms");
    const rooms = await res.json();
    const container = document.getElementById("rooms");
    const select = document.getElementById("roomSelect");
    container.innerHTML = "";
    select.innerHTML = "";

    rooms.forEach(r => {
        container.innerHTML += `
      <div class="room">
        <h3>${r.name}</h3>
        <p>📍 ${r.location}</p>
        <p>👥 Вместимость: ${r.capacity}</p>
      </div>`;
        select.innerHTML += `<option value="${r.id}">${r.name}</option>`;
    });
}

document.getElementById("bookingForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const room_id = document.getElementById("roomSelect").value;
    const start_time = document.getElementById("startTime").value;
    const end_time = document.getElementById("endTime").value;
    const purpose = document.getElementById("purpose").value;

    const res = await fetch("/bookings", {
        method: "POST",
        headers: { "Content-Type": "application/json", "X-User-Id": 1 },
        body: JSON.stringify({ room_id, user_id: 1, start_time, end_time, purpose })
    });

    const msg = document.getElementById("bookingMessage");
    if (res.status === 201) {
        msg.innerText = "Бронь создана!";
        loadRooms();
    } else {
        const text = await res.text();
        msg.innerText = "Ошибка: " + text;
    }
});

loadRooms();