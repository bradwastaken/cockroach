// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Start20_2-0]
	_ = x[GeospatialType-1]
	_ = x[AlterColumnTypeGeneral-2]
	_ = x[UserDefinedSchemas-3]
	_ = x[NoOriginFKIndexes-4]
	_ = x[NodeMembershipStatus-5]
	_ = x[MinPasswordLength-6]
	_ = x[AbortSpanBytes-7]
	_ = x[MaterializedViews-8]
	_ = x[Box2DType-9]
	_ = x[CreateLoginPrivilege-10]
	_ = x[HBAForNonTLS-11]
	_ = x[V20_2-12]
	_ = x[Start21_1-13]
	_ = x[EmptyArraysInInvertedIndexes-14]
	_ = x[UniqueWithoutIndexConstraints-15]
	_ = x[VirtualComputedColumns-16]
	_ = x[CPutInline-17]
	_ = x[ReplicaVersions-18]
	_ = x[replacedTruncatedAndRangeAppliedStateMigration-19]
	_ = x[replacedPostTruncatedAndRangeAppliedStateMigration-20]
	_ = x[NewSchemaChanger-21]
	_ = x[LongRunningMigrations-22]
	_ = x[TruncatedAndRangeAppliedStateMigration-23]
	_ = x[PostTruncatedAndRangeAppliedStateMigration-24]
	_ = x[SeparatedIntents-25]
	_ = x[TracingVerbosityIndependentSemantics-26]
	_ = x[SequencesRegclass-27]
	_ = x[ImplicitColumnPartitioning-28]
	_ = x[MultiRegionFeatures-29]
	_ = x[ClosedTimestampsRaftTransport-30]
	_ = x[ChangefeedsSupportPrimaryIndexChanges-31]
	_ = x[ForeignKeyRepresentationMigration-32]
	_ = x[PriorReadSummaries-33]
	_ = x[NonVotingReplicas-34]
	_ = x[ProtectedTsMetaPrivilegesMigration-35]
	_ = x[V21_1-36]
	_ = x[Start21_1PLUS-37]
	_ = x[Start21_2-38]
	_ = x[JoinTokensTable-39]
	_ = x[AcquisitionTypeInLeaseHistory-40]
	_ = x[SerializeViewUDTs-41]
	_ = x[ExpressionIndexes-42]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-43]
}

const _Key_name = "Start20_2GeospatialTypeAlterColumnTypeGeneralUserDefinedSchemasNoOriginFKIndexesNodeMembershipStatusMinPasswordLengthAbortSpanBytesMaterializedViewsBox2DTypeCreateLoginPrivilegeHBAForNonTLSV20_2Start21_1EmptyArraysInInvertedIndexesUniqueWithoutIndexConstraintsVirtualComputedColumnsCPutInlineReplicaVersionsreplacedTruncatedAndRangeAppliedStateMigrationreplacedPostTruncatedAndRangeAppliedStateMigrationNewSchemaChangerLongRunningMigrationsTruncatedAndRangeAppliedStateMigrationPostTruncatedAndRangeAppliedStateMigrationSeparatedIntentsTracingVerbosityIndependentSemanticsSequencesRegclassImplicitColumnPartitioningMultiRegionFeaturesClosedTimestampsRaftTransportChangefeedsSupportPrimaryIndexChangesForeignKeyRepresentationMigrationPriorReadSummariesNonVotingReplicasProtectedTsMetaPrivilegesMigrationV21_1Start21_1PLUSStart21_2JoinTokensTableAcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigration"

var _Key_index = [...]uint16{0, 9, 23, 45, 63, 80, 100, 117, 131, 148, 157, 177, 189, 194, 203, 231, 260, 282, 292, 307, 353, 403, 419, 440, 478, 520, 536, 572, 589, 615, 634, 663, 700, 733, 751, 768, 802, 807, 820, 829, 844, 873, 890, 907, 956}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
