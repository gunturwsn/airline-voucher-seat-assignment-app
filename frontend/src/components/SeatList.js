import React from "react";

function SeatList({ seats }) {
  if (seats.length === 0) return null;

  return (
    <div>
      <h3>Assigned Seats:</h3>
      <ul>
        {seats.map((seat, index) => (
          <li key={index}>{seat}</li>
        ))}
      </ul>
    </div>
  );
}

export default SeatList;
