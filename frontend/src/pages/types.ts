// types.ts
export interface Application {
    id: string;
    path: string;
    name: string;
    version: string;
  }
  
  export interface OSInfo {
    id: number;
    platform: string;
    version: string;
    build: string;
    osquery_version: string;
    timestamp: string; // ISO string
  }
  
  export interface LatestData {
    os_info: OSInfo;
    applications: Application[];
  }
  