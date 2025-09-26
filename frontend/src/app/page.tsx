import Image from "next/image";
import Link from "next/link";
import 'leaflet/dist/leaflet.css'
import Map from "../../components/map";

export default function Home() {
  return (
    <div>
      <Map />
    </div>
  );
}
