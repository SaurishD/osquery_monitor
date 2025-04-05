// app/page.tsx
'use client';

import { useEffect, useState } from 'react';
import { Application, OSInfo, LatestData } from './types';

export default function HomePage() {
  const [osInfo, setOsInfo] = useState<OSInfo | null>(null);
  const [apps, setApps] = useState<Application[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('http://localhost:8080/latest_data')
      .then(res => res.json())
      .then((data: LatestData) => {
        setOsInfo(data.os_info);
        setApps(data.applications);
        setLoading(false);
      })
      .catch(console.error);
  }, []);

  if (loading) return <p>Loading...</p>;

  return (
    <main className="p-6">
      <h1 className="text-xl font-bold mb-4">OS Information</h1>
      <div className="mb-6">
        <p><strong>Platform:</strong> {osInfo?.platform}</p>
        <p><strong>Version:</strong> {osInfo?.version}</p>
        <p><strong>Build:</strong> {osInfo?.build}</p>
        <p><strong>Osquery Version:</strong> {osInfo?.osquery_version}</p>
        {/* <p><strong>Timestamp:</strong> {new Date(osInfo?.timestamp).toLocaleString()}</p> */}
      </div>

      <h2 className="text-lg font-semibold mb-2">Applications</h2>
      <table className="min-w-full border border-collapse border-gray-300">
        <thead className="bg-gray-100">
          <tr>
            <th className="border px-4 py-2">ID</th>
            <th className="border px-4 py-2">Path</th>
            <th className="border px-4 py-2">Name</th>
            <th className="border px-4 py-2">Version</th>
          </tr>
        </thead>
        <tbody>
          {apps.map(app => (
            <tr key={`${app.id}-${app.path}`} className="hover:bg-gray-50">
              <td className="border px-4 py-2">{app.id}</td>
              <td className="border px-4 py-2">{app.path}</td>
              <td className="border px-4 py-2">{app.name}</td>
              <td className="border px-4 py-2">{app.version}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </main>
  );
}
