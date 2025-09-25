"use client";
import { MapContainer, TileLayer } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import type { LatLngExpression } from "leaflet";

const Map = () => {
  const position: LatLngExpression = [-21.114533, 55.532062];

  return (
    <div className="flex justify-center items-center">
      <MapContainer
        center={position}
        zoom={11}
        scrollWheelZoom={true}
        className="w-[980px] h-[800px] rounded-lg shadow-md"
        >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://osm.org/copyright">OpenStreetMap</a>'
        />
      </MapContainer>
    </div>
  );
};

export default Map;