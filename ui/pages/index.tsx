import * as React from "react";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import ProTip from "../src/ProTip";
import Link from "../src/Link";
import BasicTable from "../components/trackingTable";
import { ClosedIssueInfo } from "../components/types";

export async function getStaticProps() {
  const infos = (await (
    await fetch(
      "https://raw.githubusercontent.com/ichn-hu/issue-tracker/master/infos.json"
    )
  ).json()) as ClosedIssueInfo[];
  console.log(infos.length);
  return {
    props: { infos },
    revalidate: 15,
  };
}

export default function Index({ infos }: { infos: ClosedIssueInfo[] }) {
  console.log(infos[0]);
  return (
    <Container maxWidth="xl">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Issue Centric Release Management
        </Typography>
        <Link href="/about" color="secondary">
          Go to the about page
        </Link>
        <ProTip />
        <BasicTable infos={infos} />
      </Box>
    </Container>
  );
}
