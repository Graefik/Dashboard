import { toTypedSchema } from "@vee-validate/zod";
import * as zod from "zod";

export default toTypedSchema(
  zod.object({
    username: zod
      .string({ message: "Identifiant requis" })
      .min(1, { message: "Identifiant requis" })
      .max(100, { message: "Identifiant trop long" }),
    password: zod
      .string({ message: "Mot de passe requis" })
      .min(1, { message: "Mot de passe requis" }),
  }),
);
