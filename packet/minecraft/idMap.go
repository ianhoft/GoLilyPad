package minecraft

type IdMap struct {
	PacketClientKeepalive                 int
	PacketClientJoinGame                  int
	PacketClientChat                      int
	PacketClientTimeUpdate                int
	PacketClientEntityEquipment           int
	PacketClientSpawnPosition             int
	PacketClientUpdateHealth              int
	PacketClientRespawn                   int
	PacketClientPlayerPositionandLook     int
	PacketClientHeldItemChange            int
	PacketClientUseBed                    int
	PacketClientAnimation                 int
	PacketClientSpawnPlayer               int
	PacketClientCollectItem               int
	PacketClientSpawnObject               int
	PacketClientSpawnMob                  int
	PacketClientSpawnPainting             int
	PacketClientSpawnExperienceOrb        int
	PacketClientEntityVelocity            int
	PacketClientDestroyEntities           int
	PacketClientEntity                    int
	PacketClientEntityRelativeMove        int
	PacketClientEntityLook                int
	PacketClientEntityLookandRelativeMove int
	PacketClientEntityTeleport            int
	PacketClientEntityHeadLook            int
	PacketClientEntityStatus              int
	PacketClientAttachEntity              int
	PacketClientEntityMetadata            int
	PacketClientEntityEffect              int
	PacketClientRemoveEntityEffect        int
	PacketClientSetExperience             int
	PacketClientEntityProperties          int
	PacketClientChunkData                 int
	PacketClientMultiBlockChange          int
	PacketClientBlockChange               int
	PacketClientBlockAction               int
	PacketClientBlockBreakAnimation       int
	PacketClientMapChunkBulk              int
	PacketClientExplosion                 int
	PacketClientEffect                    int
	PacketClientNamedSoundEffect          int
	PacketClientParticle                  int
	PacketClientChangeGameState           int
	PacketClientSpawnGlobalEntity         int
	PacketClientOpenWindow                int
	PacketClientCloseWindow               int
	PacketClientSetSlot                   int
	PacketClientWindowItems               int
	PacketClientWindowProperty            int
	PacketClientConfirmTransaction        int
	PacketClientUpdateSign                int
	PacketClientMaps                      int
	PacketClientUpdateBlockEntity         int
	PacketClientSignEditorOpen            int
	PacketClientStatistics                int
	PacketClientPlayerList                int
	PacketClientPlayerAbilities           int
	PacketClientTabComplete               int
	PacketClientScoreboardObjective       int
	PacketClientUpdateScore               int
	PacketClientDisplayScoreboard         int
	PacketClientTeams                     int
	PacketClientPluginMessage             int
	PacketClientDisconnect                int
	PacketClientDifficulty                int
	PacketClientCombatEvent               int
	PacketClientCamera                    int
	PacketClientWorldBorder               int
	PacketClientTitle                     int
	PacketClientSetCompression            int
	PacketClientPlayerListHeadFoot        int
	PacketClientResourcePack              int
	PacketClientUpdateEntityNbt           int
	PacketServerKeepalive                 int
	PacketServerChat                      int
	PacketServerUseEntity                 int
	PacketServerPlayer                    int
	PacketServerPlayerPosition            int
	PacketServerPlayerLook                int
	PacketServerPlayerLookandPosition     int
	PacketServerPlayerDigging             int
	PacketServerPlayerBlockPlacement      int
	PacketServerHeldItemChange            int
	PacketServerAnimation                 int
	PacketServerEntityAction              int
	PacketServerSteerVehicle              int
	PacketServerCloseWindow               int
	PacketServerClickWindow               int
	PacketServerConfirmTransaction        int
	PacketServerCreativeInventoryAction   int
	PacketServerEnchantItem               int
	PacketServerUpdateSign                int
	PacketServerPlayerAbilities           int
	PacketServerTabComplete               int
	PacketServerClientSettings            int
	PacketServerClientStatus              int
	PacketServerPluginMessage             int
	PacketServerSpectate                  int
	PacketServerResourcePackStatus        int

	PacketClientLoginDisconnect      int
	PacketClientLoginEncryptRequest  int
	PacketClientLoginSuccess         int
	PacketClientLoginSetCompression  int
	PacketServerLoginStart           int
	PacketServerLoginEncryptResponse int

	// 1.9
	PacketClientBossBar         int
	PacketClientSetCooldown     int
	PacketClientUnloadChunk     int
	PacketClientVehicleMove     int
	PacketClientSetPassengers   int
	PacketServerTeleportConfirm int
	PacketServerVehicleMove     int
	PacketServerSteerBoat       int
	PacketServerUseItem         int
}
