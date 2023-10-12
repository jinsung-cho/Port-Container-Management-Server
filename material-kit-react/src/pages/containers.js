import { useCallback, useMemo, useState, useEffect } from 'react';
import Head from 'next/head';
import { Box, Container, Stack, Typography, Grid, Divider } from '@mui/material';
import { Layout as DashboardLayout } from 'src/layouts/dashboard/layout';
import { PreInformation } from 'src/sections/container/pre-information';
import { ContainerSpec } from 'src/sections/container/container-spec';
import axiosInstance from "src/axiosInstance";

const Page = () => {
  const [preInfoData, setPreInfoData] = useState([]);
  const [specData, setSpecData] = useState([]);
  useEffect(() => {
    axiosInstance.get("/preinformation")
      .then((response) => setPreInfoData(response.data))
      .catch((error) => console.error('Error fetching data: ', error));
    axiosInstance.get("/containerspec")
      .then((response) => setSpecData(response.data))
      .catch((error) => console.error('Error fetching data: ', error));
  }, []);
  return (
    <>
      <Head>
        <title>Containers</title>
      </Head>
      <Box component="main" sx={{ flexGrow: 1, py: 4 }}>
        <Container maxWidth="100%">
          <Grid container spacing={3} direction="column">
            <Grid item xs={12} sx={{ height: '45vh', minHeight: '45vh' }}>
              <Stack spacing={3}>
                <Stack direction="row" justifyContent="space-between" spacing={4}>
                  <Stack spacing={1}>
                    <Typography variant="h4">Pre-Information</Typography>
                  </Stack>
                </Stack>
                <PreInformation items={preInfoData} />
              </Stack>
            </Grid>
            <Divider sx={{ my: '10px', height: '5px' }} />
            <Grid item xs={12} sx={{ height: '45vh', minHeight: '45vh' }}>
              <Stack spacing={3}>
                <Stack direction="row" justifyContent="space-between" spacing={4}>
                  <Stack spacing={1}>
                    <Typography variant="h4">Container Spec</Typography>
                  </Stack>
                </Stack>
                <ContainerSpec items={specData} />
              </Stack>
            </Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
};

Page.getLayout = (page) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);

export default Page;
