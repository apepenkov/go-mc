// Code generated by "stringer -type ServerboundPacketID"; DO NOT EDIT.

package packetid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerboundAcceptTeleportation-0]
	_ = x[ServerboundBlockEntityTagQuery-1]
	_ = x[ServerboundChangeDifficulty-2]
	_ = x[ServerboundChatAck-3]
	_ = x[ServerboundChatCommand-4]
	_ = x[ServerboundChat-5]
	_ = x[ServerboundChatPreview-6]
	_ = x[ServerboundClientCommand-7]
	_ = x[ServerboundClientInformation-8]
	_ = x[ServerboundCommandSuggestion-9]
	_ = x[ServerboundContainerButtonClick-10]
	_ = x[ServerboundContainerClick-11]
	_ = x[ServerboundContainerClose-12]
	_ = x[ServerboundCustomPayload-13]
	_ = x[ServerboundEditBook-14]
	_ = x[ServerboundEntityTagQuery-15]
	_ = x[ServerboundInteract-16]
	_ = x[ServerboundJigsawGenerate-17]
	_ = x[ServerboundKeepAlive-18]
	_ = x[ServerboundLockDifficulty-19]
	_ = x[ServerboundMovePlayerPos-20]
	_ = x[ServerboundMovePlayerPosRot-21]
	_ = x[ServerboundMovePlayerRot-22]
	_ = x[ServerboundMovePlayerStatusOnly-23]
	_ = x[ServerboundMoveVehicle-24]
	_ = x[ServerboundPaddleBoat-25]
	_ = x[ServerboundPickItem-26]
	_ = x[ServerboundPlaceRecipe-27]
	_ = x[ServerboundPlayerAbilities-28]
	_ = x[ServerboundPlayerAction-29]
	_ = x[ServerboundPlayerCommand-30]
	_ = x[ServerboundPlayerInput-31]
	_ = x[ServerboundPong-32]
	_ = x[ServerboundRecipeBookChangeSettings-33]
	_ = x[ServerboundRecipeBookSeenRecipe-34]
	_ = x[ServerboundRenameItem-35]
	_ = x[ServerboundResourcePack-36]
	_ = x[ServerboundSeenAdvancements-37]
	_ = x[ServerboundSelectTrade-38]
	_ = x[ServerboundSetBeacon-39]
	_ = x[ServerboundSetCarriedItem-40]
	_ = x[ServerboundSetCommandBlock-41]
	_ = x[ServerboundSetCommandMinecart-42]
	_ = x[ServerboundSetCreativeModeSlot-43]
	_ = x[ServerboundSetJigsawBlock-44]
	_ = x[ServerboundSetStructureBlock-45]
	_ = x[ServerboundSignUpdate-46]
	_ = x[ServerboundSwing-47]
	_ = x[ServerboundTeleportToEntity-48]
	_ = x[ServerboundUseItemOn-49]
	_ = x[ServerboundUseItem-50]
}

const _ServerboundPacketID_name = "ServerboundAcceptTeleportationServerboundBlockEntityTagQueryServerboundChangeDifficultyServerboundChatAckServerboundChatCommandServerboundChatServerboundChatPreviewServerboundClientCommandServerboundClientInformationServerboundCommandSuggestionServerboundContainerButtonClickServerboundContainerClickServerboundContainerCloseServerboundCustomPayloadServerboundEditBookServerboundEntityTagQueryServerboundInteractServerboundJigsawGenerateServerboundKeepAliveServerboundLockDifficultyServerboundMovePlayerPosServerboundMovePlayerPosRotServerboundMovePlayerRotServerboundMovePlayerStatusOnlyServerboundMoveVehicleServerboundPaddleBoatServerboundPickItemServerboundPlaceRecipeServerboundPlayerAbilitiesServerboundPlayerActionServerboundPlayerCommandServerboundPlayerInputServerboundPongServerboundRecipeBookChangeSettingsServerboundRecipeBookSeenRecipeServerboundRenameItemServerboundResourcePackServerboundSeenAdvancementsServerboundSelectTradeServerboundSetBeaconServerboundSetCarriedItemServerboundSetCommandBlockServerboundSetCommandMinecartServerboundSetCreativeModeSlotServerboundSetJigsawBlockServerboundSetStructureBlockServerboundSignUpdateServerboundSwingServerboundTeleportToEntityServerboundUseItemOnServerboundUseItem"

var _ServerboundPacketID_index = [...]uint16{0, 30, 60, 87, 105, 127, 142, 164, 188, 216, 244, 275, 300, 325, 349, 368, 393, 412, 437, 457, 482, 506, 533, 557, 588, 610, 631, 650, 672, 698, 721, 745, 767, 782, 817, 848, 869, 892, 919, 941, 961, 986, 1012, 1041, 1071, 1096, 1124, 1145, 1161, 1188, 1208, 1226}

func (i ServerboundPacketID) String() string {
	if i < 0 || i >= ServerboundPacketID(len(_ServerboundPacketID_index)-1) {
		return "ServerboundPacketID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ServerboundPacketID_name[_ServerboundPacketID_index[i]:_ServerboundPacketID_index[i+1]]
}
