'use client';

import { useEffect, useState } from 'react';
import { Application, OSInfo, LatestData } from './types';
import {
  Box,
  CircularProgress,
  Container,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography
} from '@mui/material';

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

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      <Typography variant="h4" gutterBottom fontWeight="bold">
        OS Information
      </Typography>
      <Box mb={4}>
        <Typography><strong>Platform:</strong> {osInfo?.platform}</Typography>
        <Typography><strong>Version:</strong> {osInfo?.version}</Typography>
        <Typography><strong>Build:</strong> {osInfo?.build}</Typography>
        <Typography><strong>Osquery Version:</strong> {osInfo?.osquery_version}</Typography>
        {/* <Typography><strong>Timestamp:</strong> {new Date(osInfo?.timestamp).toLocaleString()}</Typography> */}
      </Box>

      <Typography variant="h5" gutterBottom fontWeight="medium">
        Applications
      </Typography>

      <TableContainer
  component={Paper}
  sx={{
    maxHeight: 400,
    bgcolor: '#1e1e1e', // table background
  }}
>
  <Table stickyHeader size="small" sx={{ minWidth: 650 }}>
    <TableHead>
      <TableRow sx={{ bgcolor: '#121212' }}>
        {['ID', 'Path', 'Name', 'Version'].map((header) => (
          <TableCell
            key={header}
            sx={{ color: '#ffffff', fontWeight: 'bold', bgcolor: '#121212' }}
          >
            {header}
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
    <TableBody>
      {apps.map((app, index) => (
        <TableRow
          key={`${app.id}-${app.path}`}
          sx={{
            bgcolor: index % 2 === 0 ? '#2a2a2a' : '#1e1e1e',
            '&:hover': {
              backgroundColor: '#333333',
            },
          }}
        >
          <TableCell sx={{ color: '#e0e0e0' }}>{app.id}</TableCell>
          <TableCell sx={{ color: '#e0e0e0', wordBreak: 'break-word' }}>{app.path}</TableCell>
          <TableCell sx={{ color: '#e0e0e0' }}>{app.name}</TableCell>
          <TableCell sx={{ color: '#e0e0e0' }}>{app.version}</TableCell>
        </TableRow>
      ))}
    </TableBody>
  </Table>
</TableContainer>


    </Container>
  );
}
