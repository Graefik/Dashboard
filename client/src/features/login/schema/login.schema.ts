import { toTypedSchema } from "@vee-validate/zod";
import * as zod from "zod";

export default toTypedSchema(
  zod.object({
    email: zod
      .string({ message: "Email est requis" })
      .email({ message: "Doit être un email valide" }),
    password: zod
      .string({ message: "Mot de passe est requis" })
      .min(8, { message: "Trop court" }),
  })
);
