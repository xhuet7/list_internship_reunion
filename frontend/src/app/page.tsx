import Image from "next/image";
import Link from "next/link";
import 'leaflet/dist/leaflet.css'
import Map from "../../components/map";
import Footer from "../../components/footer";
import Navbar from "../../components/navbar";

export default function Home() {
  return (
    <div>
      <Navbar />
      <Map />
      <Footer />
    </div>
  );
}
