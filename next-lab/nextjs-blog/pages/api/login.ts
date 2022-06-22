// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from "next";

type Data = {
  token: string;
};

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  console.log(req.session);
  const token = "1234";
  if (req.session) {
    Object.defineProperty(req, "session", { token });
  }

//   res.setHeader("token", token);
  res.status(200).json({ token });
}
